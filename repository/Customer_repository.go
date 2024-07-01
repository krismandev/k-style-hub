package repository

import (
	"k-style-test/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(tx *gorm.DB, customer *model.Customer) error
	GetCustomer(customers *[]model.Customer, params map[string]interface{}) error
	GetCustomerById(customer *model.Customer) error
	GetCustomerByUserID(customer *model.Customer) error
	UpdateCustomer(tx *gorm.DB, customer *model.Customer) error
	DeleteCustomer(tx *gorm.DB, customer *model.Customer) error
	CustomerCount(count *int64) error
}
