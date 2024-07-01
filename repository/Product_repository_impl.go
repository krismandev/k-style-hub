package repository

import (
	"k-style-test/model"
	"k-style-test/utility"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (r *ProductRepositoryImpl) GetProduct(products *[]model.Product, params map[string]interface{}) error {
	var err error

	result := r.DB.Scopes(utility.Paginate(params)).Scopes(utility.Order(params)).Preload("Stock").Preload("StockTransaction").Find(&products)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err

}

func (r *ProductRepositoryImpl) GetProductById(product *model.Product) error {
	var err error

	result := r.DB.Preload("Stock").Preload("StockTransaction").First(&product)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}

func (r *ProductRepositoryImpl) GetProductByProductIdList(products *[]model.Product, productIds []int) error {
	var err error

	result := r.DB.Preload("Stock").Find(&products, productIds)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}
