package usecase

import (
	"context"
	"k-style-test/model"
	"k-style-test/model/request"
	"k-style-test/model/response"
	"k-style-test/repository"
	"k-style-test/utility"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type AuthUseCaseImpl struct {
	UserRepository     repository.UserRepository
	CustomerRepository repository.CustomerRepository
	AuthLogRepository  repository.AuthLogRepository
	Validate           *validator.Validate
}

func NewAuthUseCase(validate *validator.Validate, userRepository repository.UserRepository, customerRepository repository.CustomerRepository, authLogRepository repository.AuthLogRepository) AuthUseCase {
	return &AuthUseCaseImpl{
		Validate:           validate,
		CustomerRepository: customerRepository,
		AuthLogRepository:  authLogRepository,
		UserRepository:     userRepository,
	}
}

func (usecase *AuthUseCaseImpl) Login(ctx context.Context, request request.LoginRequest) (response.LoginResponse, error) {
	var err error
	var response response.LoginResponse

	user := new(model.User)
	user.Email = request.Email

	err = usecase.Validate.Struct(request)
	if err != nil {
		logrus.Errorf("Error in UseCase : Validation Error %v", err.Error())
		return response, &utility.BadRequestError{Code: 400, Message: err.Error()}
	}

	err = usecase.UserRepository.GetUser(user)
	if err != nil {
		return response, &utility.UnauthorizedError{Code: 401, Message: "Invalid email address"}
	}

	customer := new(model.Customer)
	customer.UserID = user.ID

	err = usecase.CustomerRepository.GetCustomerByUserID(customer)
	if err != nil {
		return response, err
	}

	checkPassword := utility.ComparePass([]byte(user.Password), []byte(request.Password))
	if !checkPassword {
		return response, &utility.UnauthorizedError{Code: 401, Message: "Invalid email address or password"}
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	authLog := model.AuthLog{}
	authLog.UserID = user.ID
	authLog.LoggedInAt = &nowStr

	err = usecase.AuthLogRepository.CreateAuthLog(&authLog)
	if err != nil {
		return response, err
	}

	response.Email = request.Email
	response.UserID = user.ID
	response.CustomerID = customer.ID

	return response, err
}

func (usecase *AuthUseCaseImpl) GetAuthLog(ctx context.Context, request request.AuthLogRequest) ([]response.AuthLogResponse, int64, error) {
	var resp []response.AuthLogResponse
	var err error
	var total int64

	err = usecase.Validate.Struct(request)
	if err != nil {
		logrus.Errorf("Error in UseCase : Validation Error %v", err.Error())
		return resp, total, &utility.BadRequestError{Code: 400, Message: err.Error()}
	}

	var authLogs []model.AuthLog
	dataParams := utility.PreparePaginationAndOrderParam(request.Param)
	dataParams["user_id"] = request.UserID
	dataParams["email"] = request.Email
	dataParams["name"] = request.Name

	err = usecase.AuthLogRepository.GetAuthLog(&authLogs, dataParams)
	if err != nil {
		return resp, total, err
	}

	err = usecase.AuthLogRepository.AuthLogCount(&total)
	if err != nil {
		return resp, total, err
	}

	for _, each := range authLogs {
		data := response.ToAuthLogResponse(&each)
		resp = append(resp, data)
	}

	return resp, total, err
}
