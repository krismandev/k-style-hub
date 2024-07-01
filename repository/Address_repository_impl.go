package repository

import (
	"k-style-test/model"
	"k-style-test/utility"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressRepositoryImpl struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &AddressRepositoryImpl{DB: db}
}

func (r *AddressRepositoryImpl) GetCustomerAddress(address *model.Address) error {
	var err error
	result := r.DB.Where("customer_id = ?", address.CustomerID).Where("id = ?", address.ID).First(address)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}
