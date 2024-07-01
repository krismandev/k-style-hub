package usecase

import (
	"context"
	model "k-style-test/model"
	"k-style-test/model/request"
	"k-style-test/model/response"
	"k-style-test/repository"
	"k-style-test/utility"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerUseCaseImpl struct {
	CustomerRepository repository.CustomerRepository
	UserRepository     repository.UserRepository
	Validate           *validator.Validate
	DB                 *gorm.DB
}

func NewCustomerUseCase(customerRepository repository.CustomerRepository, userRepository repository.UserRepository, validate *validator.Validate, db *gorm.DB) CustomerUseCase {
	return &CustomerUseCaseImpl{
		CustomerRepository: customerRepository,
		Validate:           validate,
		UserRepository:     userRepository,
		DB:                 db,
	}
}

func (usecase *CustomerUseCaseImpl) RegisterCustomer(ctx context.Context, request request.RegisterCustomerRequest) (response.RegisterCustomerResponse, error) {
	var err error
	var resp response.RegisterCustomerResponse

	err = usecase.Validate.Struct(request)
	if err != nil {
		logrus.Errorf("Error in UseCase : Validation Error %v", err.Error())
		return resp, &utility.BadRequestError{Code: 400, Message: err.Error()}
	}

	user := new(model.User)
	user.Email = request.Email
	user.Password = utility.HashPassword(request.Password)

	tx := usecase.DB.Begin()

	err = usecase.UserRepository.CreateUser(tx, user)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	customer := new(model.Customer)

	customer.FirstName = request.FirstName
	customer.LastName = utility.NullableString(request.LastName)
	customer.Gender = request.Gender
	customer.PhoneNumber = utility.NullableString(request.PhoneNumber)
	customer.UserID = user.ID
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	customer.CreatedAt = &nowStr

	err = usecase.CustomerRepository.CreateCustomer(tx, customer)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	tx.Commit()
	resp = response.ToCustomerResponse(customer)

	return resp, err
}

func (usecase *CustomerUseCaseImpl) GetCustomer(ctx context.Context, request request.GetCustomerRequest) ([]response.GetCustomerResponse, int64, error) {
	var resp []response.GetCustomerResponse
	var err error
	var total int64

	var customerList []model.Customer
	dataParams := utility.PreparePaginationAndOrderParam(request.Param)
	dataParams["name"] = request.Name
	dataParams["email"] = request.Email

	err = usecase.CustomerRepository.GetCustomer(&customerList, dataParams)
	if err != nil {
		return resp, total, err
	}

	var responses []response.GetCustomerResponse
	for _, each := range customerList {
		dt := response.ToGetCustomerResponse(&each)
		responses = append(responses, dt)
	}

	err = usecase.CustomerRepository.CustomerCount(&total)

	return responses, total, err
}

func (usecase *CustomerUseCaseImpl) UpdateCustomer(ctx context.Context, request request.UpdateCustomerRequest) (response.UpdateCustomerResponse, error) {
	var resp response.UpdateCustomerResponse
	var err error

	err = usecase.Validate.Struct(request)
	if err != nil {
		logrus.Errorf("Error in UseCase : Validation Error %v", err.Error())
		return resp, &utility.BadRequestError{Code: 400, Message: err.Error()}
	}
	existingCustomer := new(model.Customer)
	existingCustomer.ID = request.CustomerID
	err = usecase.CustomerRepository.GetCustomerById(existingCustomer)
	if err != nil {
		logrus.Errorf("Error in UseCase : Customer Not Found %v", err)
		return resp, &utility.NotFoundError{Code: 404, Message: err.Error()}
	}

	existingCustomer.FirstName = request.FirstName
	existingCustomer.LastName = utility.NullableString(request.LastName)
	existingCustomer.PhoneNumber = utility.NullableString(request.PhoneNumber)
	existingCustomer.Gender = request.Gender

	tx := usecase.DB.Begin()
	err = usecase.CustomerRepository.UpdateCustomer(tx, existingCustomer)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	tx.Commit()

	resp = response.ToUpdateCustomerResponse(existingCustomer)

	return resp, err
}

func (usecase *CustomerUseCaseImpl) DeleteCustomer(ctx context.Context, request request.DeleteCustomerRequest) (response.DeleteCustomerResponse, error) {
	var resp response.DeleteCustomerResponse
	var err error

	err = usecase.Validate.Struct(request)
	if err != nil {
		logrus.Errorf("Error in UseCase : Validation Error %v", err.Error())
		return resp, &utility.BadRequestError{Code: 400, Message: err.Error()}
	}

	existingCustomer := new(model.Customer)
	existingCustomer.ID = request.CustomerID
	err = usecase.CustomerRepository.GetCustomerById(existingCustomer)
	if err != nil {
		logrus.Errorf("Error in UseCase : Customer Not Found %v", err)
		return resp, &utility.NotFoundError{Code: 404, Message: err.Error()}
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	existingCustomer.DeletedAt = &nowStr
	tx := usecase.DB.Begin()
	err = usecase.CustomerRepository.DeleteCustomer(tx, existingCustomer)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	tx.Commit()

	resp = response.ToDeleteCustomerResponse(existingCustomer)

	return resp, err
}
