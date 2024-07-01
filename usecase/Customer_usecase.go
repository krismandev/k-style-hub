package usecase

import (
	"context"
	"k-style-test/model/request"
	"k-style-test/model/response"
)

type CustomerUseCase interface {
	// GetUser(ctx context.Context, request request.UserRequest) (response.UserResponse, error)
	// GetUsers(ctx context.Context, request request.UserRequest) ([]response.UserResponse, error)
	RegisterCustomer(ctx context.Context, request request.RegisterCustomerRequest) (response.RegisterCustomerResponse, error)
	GetCustomer(ctx context.Context, request request.GetCustomerRequest) ([]response.GetCustomerResponse, int64, error)
	UpdateCustomer(ctx context.Context, request request.UpdateCustomerRequest) (response.UpdateCustomerResponse, error)
	DeleteCustomer(ctx context.Context, request request.DeleteCustomerRequest) (response.DeleteCustomerResponse, error)
}
