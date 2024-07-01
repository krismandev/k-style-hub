package usecase

import (
	"context"
	"errors"
	"k-style-test/model"
	"k-style-test/model/request"
	"k-style-test/model/response"
	"k-style-test/repository"
	"k-style-test/utility"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderUseCaseImpl struct {
	DB                 *gorm.DB
	OrderRepository    repository.OrderRepository
	CustomerRepository repository.CustomerRepository
	ProductRepository  repository.ProductRepository
	StockRepository    repository.StockRepository
	AddressRepository  repository.AddressRepository
	Validate           *validator.Validate
}

func NewOrderUseCase(db *gorm.DB, orderRepository repository.OrderRepository, validate *validator.Validate, customerRepository repository.CustomerRepository, productRepository repository.ProductRepository, stockRepository repository.StockRepository, addressRepository repository.AddressRepository) OrderUseCase {
	return &OrderUseCaseImpl{
		DB:                 db,
		OrderRepository:    orderRepository,
		CustomerRepository: customerRepository,
		ProductRepository:  productRepository,
		StockRepository:    stockRepository,
		AddressRepository:  addressRepository,
		Validate:           validate,
	}
}

func (usecase *OrderUseCaseImpl) GetOrder(ctx context.Context, request request.GetOrderRequest) ([]response.OrderResponse, int64, error) {
	var responses []response.OrderResponse
	var err error
	var total int64 = 0

	err = usecase.Validate.Struct(request)
	if err != nil {
		logrus.Errorf("Error in UseCase : Validation Error %v", err.Error())
		return responses, total, &utility.BadRequestError{Code: 400, Message: err.Error()}
	}

	var orders []model.Order
	dataParams := utility.PreparePaginationAndOrderParam(request.Param)
	dataParams["customer_id"] = request.CustomerID
	dataParams["code"] = request.Code

	err = usecase.OrderRepository.GetOrder(&orders, dataParams)
	if err != nil {
		return responses, total, err
	}

	// var responses []response.GetCustomerResponse
	for _, each := range orders {
		dt := response.ToOrderResponse(&each)
		responses = append(responses, dt)
	}

	err = usecase.OrderRepository.OrderCount(request.CustomerID, &total)
	if err != nil {
		return responses, total, err
	}

	return responses, total, err
}

func (usecase *OrderUseCaseImpl) GetDetailOrder(ctx context.Context, request request.GetOrderRequest) (response.OrderResponse, error) {
	var resp response.OrderResponse
	var err error

	err = usecase.Validate.Struct(request)
	if err != nil {
		logrus.Errorf("Error in UseCase : Validation Error %v", err.Error())
		return resp, &utility.BadRequestError{Code: 400, Message: err.Error()}
	}

	order := new(model.Order)
	order.ID = request.OrderID
	order.CustomerID = request.CustomerID

	err = usecase.OrderRepository.GetOrderByID(order)
	if err != nil {
		return resp, err
	}

	resp = response.ToOrderDetailResponse(order)

	return resp, err
}

func (usecase *OrderUseCaseImpl) CreateOrder(ctx context.Context, request request.CreateOrderRequest) (response.CreateOrderResponse, error) {
	var resp response.CreateOrderResponse
	var err error

	err = usecase.Validate.Struct(request)

	if err != nil {
		logrus.Errorf("Error in UseCase : Validation Error %v", err.Error())
		return resp, &utility.BadRequestError{Message: err.Error()}
	}

	// customer := new(model.Customer)
	// customer.UserID = request.UserID
	var address model.Address
	address.CustomerID = request.CustomerID
	address.ID = request.AddressID
	err = usecase.AddressRepository.GetCustomerAddress(&address)
	if err != nil {
		return resp, &utility.UnprocessableContentError{Message: "Invalid address"}
	}

	tx := usecase.DB.Begin()
	order := new(model.Order)

	order.CustomerID = request.CustomerID
	order.AddressID = request.AddressID
	order.Code = utility.GenerateOrderCode()
	order.OrderStatus = 1
	strNow := time.Now().Format("2006-01-02 15:04:05")
	order.CreatedAt = utility.NullableString(strNow)

	err = usecase.OrderRepository.CreateOrder(tx, order)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	var products []model.Product
	var listProductId []int
	for _, each := range request.OrderDetails {
		listProductId = append(listProductId, each.ProductID)
	}

	err = usecase.ProductRepository.GetProductByProductIdList(&products, listProductId)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	var stocks []model.Stock
	err = usecase.StockRepository.GetProductStockByProductIdListTransaction(tx, &stocks, listProductId)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	var details []model.OrderDetail
	var grandTotal float64
	grandTotal, details, err = usecase.ProcessOrderDetail(products, stocks, order.ID, request.OrderDetails)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	err = usecase.OrderRepository.CreateOrderDetail(tx, &details)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	var stockTransactions []model.StockTransaction

	stockTransactions, err = usecase.ProcessStockTransaction(tx, details, stocks, strNow, 2)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	err = usecase.StockRepository.CreateStockTransaction(tx, &stockTransactions)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	order.GrandTotal = grandTotal

	err = usecase.OrderRepository.UpdateGrandTotal(tx, order.ID, grandTotal)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	tx.Commit()

	resp = response.ToCreateOrderResponse(order)

	return resp, err
}

func (usecase *OrderUseCaseImpl) UpdateOrder(ctx context.Context, req request.UpdateOrderRequest) (response.UpdateOrderResponse, error) {
	var resp response.UpdateOrderResponse
	var err error

	err = usecase.Validate.Struct(req)

	if err != nil {
		logrus.Errorf("Error in UseCase : Validation Error %v", err.Error())
		return resp, &utility.BadRequestError{Message: err.Error()}
	}

	order := new(model.Order)
	order.ID = req.OrderID
	order.CustomerID = req.CustomerID

	err = usecase.OrderRepository.GetOrderByID(order)
	if err != nil {
		return resp, err
	}

	var address model.Address
	address.CustomerID = req.CustomerID
	address.ID = req.AddressID
	err = usecase.AddressRepository.GetCustomerAddress(&address)
	if err != nil {
		return resp, &utility.UnprocessableContentError{Message: "Invalid address"}
	}

	tx := usecase.DB.Begin()

	// need to be canceled first
	if order.OrderStatus != 3 {
		tx.Rollback()
		logrus.Errorf("Error in UseCase : The Order needs to be canceled first")
		return resp, &utility.UnprocessableContentError{Message: "The Order needs to be canceled first"}
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")

	var products []model.Product
	var listProductId []int
	for _, each := range req.OrderDetails {
		listProductId = append(listProductId, each.ProductID)
	}

	err = usecase.ProductRepository.GetProductByProductIdList(&products, listProductId)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	var stocks []model.Stock
	err = usecase.StockRepository.GetProductStockByProductIdListTransaction(tx, &stocks, listProductId)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	var paramOrderDetails []model.OrderDetail
	var grandTotal float64
	for _, each := range order.OrderDetail {
		exist := false

		for _, reqDetail := range req.OrderDetails {
			if each.ProductID == reqDetail.ProductID {
				exist = true
			}
		}

		if !exist {
			logrus.Errorf("Error in UseCase : You can only change the quantity from the previous details")
			return resp, &utility.ConflictError{Message: "You can only change the quantity from the previous details"}
		}
	}

	// lakukan sebaliknya
	for _, each := range req.OrderDetails {
		exist := false

		for _, prevDetail := range order.OrderDetail {
			if each.ProductID == prevDetail.ProductID {
				exist = true
			}
		}

		if !exist {
			logrus.Errorf("Error in UseCase : You can only change the quantity from the previous details")
			return resp, &utility.ConflictError{Message: "You can only change the quantity from the previous details"}
		}

	}

	grandTotal, paramOrderDetails, err = usecase.ProcessOrderDetail(products, stocks, order.ID, req.OrderDetails)
	if err != nil {
		tx.Rollback()
		return resp, err
	}
	order.GrandTotal = grandTotal
	order.AddressID = req.AddressID
	order.UpdatedAt = &nowStr
	order.OrderStatus = 1

	err = usecase.OrderRepository.UpdateOrder(tx, order)
	if err != nil {
		logrus.Errorf("Error in UseCase : Failed to update order %v", err)
		tx.Rollback()
		return resp, err
	}

	for _, each := range paramOrderDetails {
		err = usecase.OrderRepository.UpdateOrderDetail(tx, &each)
		if err != nil {
			tx.Rollback()
			return resp, err
		}
	}

	var stockTransactions []model.StockTransaction
	stockTransactions, err = usecase.ProcessStockTransaction(tx, paramOrderDetails, stocks, nowStr, 2)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	err = usecase.StockRepository.CreateStockTransaction(tx, &stockTransactions)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	order.OrderDetail = paramOrderDetails

	tx.Commit()

	resp = response.ToUpdateOrderResponse(order)

	return resp, err
}

func (usecase *OrderUseCaseImpl) CancelOrder(ctx context.Context, request request.CancelOrderRequest) (response.CancelOrderRespoonse, error) {
	var resp response.CancelOrderRespoonse
	var err error

	err = usecase.Validate.Struct(request)

	// defer tx.Commit()
	if err != nil {
		logrus.Errorf("Error in UseCase : Validation Error %v", err.Error())
		return resp, &utility.BadRequestError{Message: err.Error()}
	}
	tx := usecase.DB.Begin()
	nowStr := time.Now().Format("2006-01-02 15:04:05")

	order := new(model.Order)
	order.ID = request.OrderID
	order.CustomerID = request.CustomerID
	order.UpdatedAt = &nowStr
	err = usecase.OrderRepository.GetOrderByID(order)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("Error in UseCase : Order not found %v", err)
		return resp, &utility.NotFoundError{Message: "Order not found"}
	}

	// payment processed / has already canceled
	if order.OrderStatus == 2 || order.OrderStatus == 3 {
		tx.Rollback()
		logrus.Errorf("Error in UseCase : Can not cancel order")
		return resp, &utility.ConflictError{Message: "Can not cancel order"}
	}

	order.OrderStatus = 3
	order.UpdatedAt = &nowStr
	err = usecase.OrderRepository.CancelOrder(tx, order)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	var stocks []model.Stock
	var listProductId []int
	for _, each := range order.OrderDetail {
		listProductId = append(listProductId, each.ProductID)
	}
	err = usecase.StockRepository.GetProductStockByProductIdListTransaction(tx, &stocks, listProductId)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	var stockTransactions []model.StockTransaction

	stockTransactions, err = usecase.ProcessStockTransaction(tx, order.OrderDetail, stocks, nowStr, 3)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	err = usecase.StockRepository.CreateStockTransaction(tx, &stockTransactions)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	resp = response.ToCancelOrderResponse(order)

	tx.Commit()

	return resp, err
}

func (usecase *OrderUseCaseImpl) ValidateStock(orderDetail request.OrderDetailRequest, stocks []model.Stock) error {
	var err error

	for _, stock := range stocks {
		if stock.ProductID == orderDetail.ProductID {
			if orderDetail.Quantity > stock.Quantity {
				return errors.New("Requested quantity exceeds available stock")
			}
		}
	}

	return err
}

func (usecase *OrderUseCaseImpl) ProcessOrderDetail(products []model.Product, stocks []model.Stock, orderID int, orderDetailsReq []request.OrderDetailRequest) (gTotal float64, orderDetails []model.OrderDetail, errorProcess error) {
	var err error

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	var grandTotal float64
	var details []model.OrderDetail
	for _, each := range orderDetailsReq {
		detail := model.OrderDetail{}
		detail.OrderID = orderID
		detail.ProductID = each.ProductID
		detail.Quantity = each.Quantity
		var productPrice float64
		productFound := false
		for _, prd := range products {
			if prd.ID == each.ProductID {
				productPrice = prd.Price
				productFound = true
				break
			}
		}

		if !productFound {
			logrus.Errorf("Error in UseCase : Invalid Product")
			return grandTotal, details, &utility.UnprocessableContentError{Message: "Invalid Product"}
		}

		err = usecase.ValidateStock(each, stocks)
		if err != nil {
			logrus.Errorf("Error in UseCase : %v", err)
			return grandTotal, details, &utility.UnprocessableContentError{Message: err.Error()}
		}

		amount := float64(each.Quantity) * productPrice
		detail.Amount = amount
		detail.CreatedAt = nowStr
		detail.Price = productPrice

		details = append(details, detail)
		grandTotal = grandTotal + amount
	}

	return grandTotal, details, err
}

func (usecase *OrderUseCaseImpl) ProcessStockTransaction(tx *gorm.DB, details []model.OrderDetail, stocks []model.Stock, timeNow string, transactionType int) ([]model.StockTransaction, error) {
	var err error
	var output []model.StockTransaction

	for _, detail := range details {
		stockTransaction := model.StockTransaction{}
		stockTransaction.ProductID = detail.ProductID
		stockTransaction.CreatedAt = timeNow
		var lastStock int
		for _, stock := range stocks {
			if stock.ProductID == detail.ProductID {
				lastStock = stock.Quantity
				break
			}
		}
		var qtyChange int = detail.Quantity
		if transactionType == 2 {
			qtyChange = detail.Quantity * -1
		}
		stockAfter := lastStock + qtyChange
		stockTransaction.LastStock = lastStock
		stockTransaction.StockAfter = stockAfter
		stockTransaction.QuantityChange = qtyChange
		stockTransaction.TransactionType = transactionType

		output = append(output, stockTransaction)

		err = usecase.StockRepository.UpdateProductStock(tx, detail.ProductID, stockAfter)
		if err != nil {
			return output, err
		}
	}

	return output, err
}
