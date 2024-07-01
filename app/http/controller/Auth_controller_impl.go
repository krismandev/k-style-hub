package controller

import (
	"k-style-test/app/http/middleware"
	"k-style-test/model/request"
	"k-style-test/model/response"
	"k-style-test/usecase"
	"k-style-test/utility"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AuthControllerImpl struct {
	AuthUseCase usecase.AuthUseCase
}

func NewAuthController(authUseCase usecase.AuthUseCase) AuthController {
	return &AuthControllerImpl{
		AuthUseCase: authUseCase,
	}
}

func (controller *AuthControllerImpl) Login(c echo.Context) error {
	var err error
	ctx := c.Request().Context()

	var responseData interface{}
	request := request.LoginRequest{}

	err = utility.ParseRequestBody(c, &request)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseSingleJSON(c, responseData, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}

	responseData, err = controller.AuthUseCase.Login(ctx, request)
	if err != nil {
		utility.WriteResponseSingleJSON(c, responseData, err)
		return err
	}

	loginResp := responseData.(response.LoginResponse)
	token, expiredAt, err := middleware.GenerateJWT(loginResp.UserID, loginResp.CustomerID)
	loginResp.Token = token
	loginResp.ExpiredAt = expiredAt

	loginResp.UserID = 0
	loginResp.CustomerID = 0

	utility.WriteResponseSingleJSON(c, loginResp, err)

	return err
}

func (controller *AuthControllerImpl) GetAuthLog(c echo.Context) error {
	var err error

	var dttb response.GlobalListDataTableResponse
	ctx := c.Request().Context()

	authLogRequest := request.AuthLogRequest{}
	err = utility.ParseRequestBody(c, &authLogRequest)
	if err != nil {
		logrus.Info("ParsingError : ", err)
		utility.WriteResponseListJSON(c, dttb, &utility.BadRequestError{Code: 400, Message: err.Error()})
		return err
	}

	authLogsData, total, err := controller.AuthUseCase.GetAuthLog(ctx, authLogRequest)
	if err != nil {
		utility.WriteResponseListJSON(c, dttb, err)
	}

	for _, each := range authLogsData {
		dttb.List = append(dttb.List, each)
	}

	dttb.TotalData = total
	dttb.PerPage = authLogRequest.Param.PerPage
	dttb.Page = authLogRequest.Param.Page

	utility.WriteResponseListJSON(c, dttb, err)

	return err
}
