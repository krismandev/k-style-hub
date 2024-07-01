package controller

import (
	"k-style-test/app/http/middleware"
	"k-style-test/model/request"
	"k-style-test/model/response"
	"k-style-test/usecase"
	"k-style-test/utility"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type OrderControllerImpl struct {
	OrderUseCase usecase.OrderUseCase
}

func NewOrderController(orderUseCase usecase.OrderUseCase) OrderController {
	return &OrderControllerImpl{
		OrderUseCase: orderUseCase,
	}
}

func (controller *OrderControllerImpl) GetOrder(c echo.Context) error {
	var err error
	var dttb response.GlobalListDataTableResponse
	ctx := c.Request().Context()
	authUser := middleware.GetAuthUser(c)

	orderRequest := request.GetOrderRequest{}
	err = utility.ParseRequestBody(c, &orderRequest)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseListJSON(c, dttb, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}
	orderRequest.CustomerID = authUser.CustomerID

	ordersData, total, err := controller.OrderUseCase.GetOrder(ctx, orderRequest)
	if err != nil {
		utility.WriteResponseListJSON(c, dttb, err)
		return err
	}

	for _, each := range ordersData {
		dttb.List = append(dttb.List, each)
	}

	dttb.TotalData = total
	dttb.PerPage = orderRequest.Param.PerPage
	dttb.Page = orderRequest.Param.Page

	utility.WriteResponseListJSON(c, dttb, err)

	return err
}

func (controller *OrderControllerImpl) GetDetailOrder(c echo.Context) error {
	var err error
	ctx := c.Request().Context()
	authUser := middleware.GetAuthUser(c)

	var responseData interface{}

	orderRequest := request.GetOrderRequest{}
	err = utility.ParseRequestBody(c, &orderRequest)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}

	orderRequest.CustomerID = authUser.CustomerID
	orderID := c.Param("id")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		logrus.Info("Convert to int Error : ", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: "Invalid ID"})
		return err
	}

	orderRequest.OrderID = orderIDInt

	responseData, err = controller.OrderUseCase.GetDetailOrder(ctx, orderRequest)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}

	utility.WriteResponseSingleJSON(c, responseData, err)

	return err
}

func (controller *OrderControllerImpl) CreateOrder(c echo.Context) error {
	var err error

	ctx := c.Request().Context()

	authUser := middleware.GetAuthUser(c)

	var responseData interface{}
	createOrderRequest := request.CreateOrderRequest{}
	err = utility.ParseRequestBody(c, &createOrderRequest)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}
	createOrderRequest.UserID = authUser.UserID
	createOrderRequest.CustomerID = authUser.CustomerID

	responseData, err = controller.OrderUseCase.CreateOrder(ctx, createOrderRequest)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return err
	}

	utility.WriteResponseSingleJSON(c, responseData, err)

	return err
}

func (controller *OrderControllerImpl) CancelOrder(c echo.Context) error {
	var err error

	ctx := c.Request().Context()

	authUser := middleware.GetAuthUser(c)

	var responseData interface{}
	cancelOrderRequest := request.CancelOrderRequest{}
	err = utility.ParseRequestBody(c, &cancelOrderRequest)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}
	cancelOrderRequest.CustomerID = authUser.CustomerID

	responseData, err = controller.OrderUseCase.CancelOrder(ctx, cancelOrderRequest)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return err
	}

	utility.WriteResponseSingleJSON(c, responseData, err)

	return err
}

func (controller *OrderControllerImpl) UpdateOrder(c echo.Context) error {
	var err error

	ctx := c.Request().Context()

	authUser := middleware.GetAuthUser(c)

	var responseData interface{}
	updateOrderRequest := request.UpdateOrderRequest{}
	err = utility.ParseRequestBody(c, &updateOrderRequest)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}
	updateOrderRequest.UserID = authUser.UserID
	updateOrderRequest.CustomerID = authUser.CustomerID

	responseData, err = controller.OrderUseCase.UpdateOrder(ctx, updateOrderRequest)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return err
	}

	utility.WriteResponseSingleJSON(c, responseData, err)

	return err
}
