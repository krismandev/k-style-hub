package repository

import "k-style-test/model"

type AddressRepository interface {
	GetCustomerAddress(address *model.Address) error
}
