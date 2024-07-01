package repository

import "k-style-test/model"

type ProductRepository interface {
	GetProduct(products *[]model.Product, params map[string]interface{}) error
	GetProductById(product *model.Product) error
	GetProductByProductIdList(products *[]model.Product, productIds []int) error
}
