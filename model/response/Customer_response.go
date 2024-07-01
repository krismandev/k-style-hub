package response

import (
	"k-style-test/model"
)

type RegisterCustomerResponse struct {
	CustomerID  int     `json:"customer_id"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name"`
	Email       string  `json:"email"`
	Gender      string  `json:"gender"`
	PhoneNumber *string `json:"phone_number"`
}

type UpdateCustomerResponse struct {
	CustomerID  int     `json:"customer_id"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name"`
	Gender      string  `json:"gender"`
	PhoneNumber *string `json:"phone_number"`
}

type DeleteCustomerResponse struct {
	DeletedAt string `json:"deleted_at"`
}

type GetCustomerResponse struct {
	ID          int     `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name"`
	Email       string  `json:"email"`
	Gender      string  `json:"gender"`
	PhoneNumber *string `json:"phone_number"`
}

func ToGetCustomerResponse(customer *model.Customer) GetCustomerResponse {
	var response GetCustomerResponse
	response.ID = customer.ID
	response.FirstName = customer.FirstName
	response.LastName = customer.LastName
	response.Email = customer.User.Email
	response.Gender = customer.Gender
	response.PhoneNumber = customer.PhoneNumber

	return response
}

func ToUpdateCustomerResponse(customer *model.Customer) UpdateCustomerResponse {
	var response UpdateCustomerResponse

	response.CustomerID = customer.ID
	response.FirstName = customer.FirstName
	response.LastName = customer.LastName
	response.Gender = customer.Gender
	response.PhoneNumber = customer.PhoneNumber

	return response
}

func ToDeleteCustomerResponse(customer *model.Customer) DeleteCustomerResponse {
	var response DeleteCustomerResponse

	response.DeletedAt = *customer.DeletedAt

	return response
}

func ToCustomerResponse(customer *model.Customer) RegisterCustomerResponse {
	var response RegisterCustomerResponse
	response.CustomerID = customer.ID
	response.FirstName = customer.FirstName
	response.LastName = customer.LastName
	response.Gender = customer.Gender
	response.PhoneNumber = customer.PhoneNumber

	return response
}
