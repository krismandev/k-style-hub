package repository

import (
	"k-style-test/model"
	"k-style-test/utility"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		DB: db,
	}
}

func (r *OrderRepositoryImpl) OrderCount(customerID int, count *int64) error {
	var err error
	result := r.DB.Model(&model.Order{}).Where("customer_id = ?", customerID).Count(count)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (r *OrderRepositoryImpl) GetOrder(orders *[]model.Order, params map[string]interface{}) error {
	var err error

	qry := r.DB.Scopes(utility.Paginate(params)).Scopes(utility.Order(params))
	if _, ok := params["customer_id"]; ok && params["customer_id"].(int) > 0 {
		qry = qry.Where("customer_id = ?", params["customer_id"])
	}
	if _, ok := params["code"]; ok && len(params["code"].(string)) > 0 {
		qry = qry.Where("code like ?", "%"+params["code"].(string)+"%")
	}
	result := qry.Find(&orders)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (r *OrderRepositoryImpl) GetOrderByID(order *model.Order) error {
	var err error

	result := r.DB.Preload("OrderDetail").Where("customer_id = ?", order.CustomerID).Where("id = ? ", order.ID).First(&order)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}

func (r *OrderRepositoryImpl) CreateOrder(tx *gorm.DB, order *model.Order) error {
	var err error

	result := tx.Select("Code", "GrandTotal", "AddressID", "CreatedAt", "CustomerID").Create(&order)
	err = utility.CheckErrorResult(result)
	if err != nil {

		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (r *OrderRepositoryImpl) CreateOrderDetail(tx *gorm.DB, orderDetails *[]model.OrderDetail) error {
	var err error

	result := tx.Select("OrderID", "ProductID", "AddressID", "Price", "Quantity", "Amount", "CreatedAt", "Status").Create(&orderDetails)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}

func (r *OrderRepositoryImpl) UpdateGrandTotal(tx *gorm.DB, orderID int, grandTotal float64) error {
	var err error

	result := tx.Model(&model.Order{}).Where("id = ?", orderID).Update("grand_total", grandTotal)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}

func (r *OrderRepositoryImpl) UpdateOrder(tx *gorm.DB, order *model.Order) error {
	var err error

	result := tx.Select("GrandTotal", "AddressID", "UpdatedAt").Updates(&order)
	err = utility.CheckErrorResult(result)
	if err != nil {

		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}

func (r *OrderRepositoryImpl) UpdateOrderDetail(tx *gorm.DB, orderDetails *model.OrderDetail) error {
	var err error

	result := tx.Select("Quantity", "Amount", "Status", "UpdatedAt").Where("product_id = ? AND order_id = ?", orderDetails.ProductID, orderDetails.OrderID).Updates(orderDetails)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}

func (r *OrderRepositoryImpl) CancelOrder(tx *gorm.DB, order *model.Order) error {
	var err error

	result := tx.Select("OrderStatus", "UpdatedAt").Where("id = ?", order.ID).Updates(order)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}
