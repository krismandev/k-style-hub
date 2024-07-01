package repository

import (
	"k-style-test/model"
	"k-style-test/utility"

	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{
		DB: db,
	}
}

func (r *CustomerRepositoryImpl) GetCustomer(customers *[]model.Customer, params map[string]interface{}) error {
	var err error

	qry := r.DB.Scopes(utility.Paginate(params)).Joins("User").Scopes(utility.Order(params))
	if _, ok := params["name"]; ok && len(params["name"].(string)) > 0 {
		qry = qry.Where("name like ?", "%"+params["name"].(string)+"%")
	}
	if _, ok := params["email"]; ok && len(params["email"].(string)) > 0 {
		qry = qry.Where("email like ?", "%"+params["email"].(string)+"%")
	}
	result := qry.Where("deleted_at IS NULL").Find(&customers)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (r *CustomerRepositoryImpl) CreateCustomer(tx *gorm.DB, customer *model.Customer) error {
	var err error

	result := tx.Select("Email", "Password", "FirstName", "LastName", "Gender", "UserID", "CreatedAt").Create(&customer)
	err = utility.CheckErrorResult(result)
	if err != nil {

		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (r *CustomerRepositoryImpl) UpdateCustomer(tx *gorm.DB, customer *model.Customer) error {
	var err error

	result := tx.Model(&customer).Select("FirstName", "LastName", "PhoneNumber", "Gender").Updates(customer)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}

func (r *CustomerRepositoryImpl) GetCustomerById(customer *model.Customer) error {
	var err error

	result := r.DB.First(&customer)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}

func (r *CustomerRepositoryImpl) DeleteCustomer(tx *gorm.DB, customer *model.Customer) error {
	var err error

	result := tx.Model(&customer).Select("DeletedAt").Updates(customer)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}

func (r *CustomerRepositoryImpl) GetCustomerByUserID(customer *model.Customer) error {
	var err error

	result := r.DB.Where("user_id = ?", customer.UserID).First(&customer)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}

func (r *CustomerRepositoryImpl) CustomerCount(count *int64) error {
	var err error
	result := r.DB.Model(&model.Customer{}).Count(count)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}
