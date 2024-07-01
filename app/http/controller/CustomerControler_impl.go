package controller

import (
	"k-style-test/model/request"
	"k-style-test/model/response"
	"k-style-test/usecase"
	"k-style-test/utility"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CustomerControllerImpl struct {
	CustomerUseCase usecase.CustomerUseCase
}

func NewCustomerController(useCase usecase.CustomerUseCase) CustomerController {
	return &CustomerControllerImpl{
		CustomerUseCase: useCase,
	}
}

func (controller *CustomerControllerImpl) Register(c echo.Context) error {
	var err error
	ctx := c.Request().Context()

	var responseData interface{}
	createUserRequest := request.RegisterCustomerRequest{}
	err = utility.ParseRequestBody(c, &createUserRequest)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}

	responseData, err = controller.CustomerUseCase.RegisterCustomer(ctx, createUserRequest)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return err
	}

	utility.WriteResponseSingleJSON(c, responseData, err)

	return err
}

// func (controller *CustomerControllerImpl) GetUsers(c echo.Context) error {
// 	var err error
// 	ctx := c.Request().Context()

// 	var responseData interface{}
// 	userReq := request.UserRequest{}
// 	err = utility.ParseRequestBody(c, &userReq)
// 	if err != nil {
// 		logrus.Info("ParsingError : ", err)
// 		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
// 		return err
// 	}

// 	responseData, err = controller.CustomerUseCase.GetUsers(ctx, userReq)
// 	if err != nil {
// 		utility.WriteResponseSingleJSON(c, responseData, err)
// 		return err
// 	}

// 	utility.WriteResponseSingleJSON(c, responseData, err)
// 	return err
// }

func (controller *CustomerControllerImpl) GetCustomer(c echo.Context) error {
	var err error
	var dttb response.GlobalListDataTableResponse

	ctx := c.Request().Context()

	customerReq := request.GetCustomerRequest{}
	err = utility.ParseRequestBody(c, &customerReq)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseListJSON(c, dttb, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}

	customerData, total, err := controller.CustomerUseCase.GetCustomer(ctx, customerReq)
	if err != nil {
		utility.WriteResponseListJSON(c, dttb, err)
		return err
	}

	for _, each := range customerData {
		dttb.List = append(dttb.List, each)
	}

	dttb.TotalData = total
	dttb.PerPage = customerReq.Param.PerPage
	dttb.Page = customerReq.Param.Page

	utility.WriteResponseListJSON(c, dttb, err)
	return err

}

func (controller *CustomerControllerImpl) UpdateCustomer(c echo.Context) error {
	var err error

	var responseData interface{}
	ctx := c.Request().Context()

	customerReq := request.UpdateCustomerRequest{}
	err = utility.ParseRequestBody(c, &customerReq)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}

	responseData, err = controller.CustomerUseCase.UpdateCustomer(ctx, customerReq)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return err
	}

	utility.WriteResponseSingleJSON(c, responseData, err)

	return err
}

func (controller *CustomerControllerImpl) DeleteCustomer(c echo.Context) error {
	var err error

	var responseData interface{}
	ctx := c.Request().Context()

	customerReq := request.DeleteCustomerRequest{}
	err = utility.ParseRequestBody(c, &customerReq)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}

	responseData, err = controller.CustomerUseCase.DeleteCustomer(ctx, customerReq)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return err
	}

	utility.WriteResponseSingleJSON(c, responseData, err)

	return err
}
