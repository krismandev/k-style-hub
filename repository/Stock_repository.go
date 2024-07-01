package repository

import (
	"k-style-test/model"

	"gorm.io/gorm"
)

type StockRepository interface {
	GetProductStockById(stock *model.Stock) error
	GetProductStockByProductIdListTransaction(tx *gorm.DB, stocks *[]model.Stock, productIds []int) error
	UpdateProductStock(tx *gorm.DB, productID int, stock int) error
	CreateStockTransaction(tx *gorm.DB, stocks *[]model.StockTransaction) error
	AddStock(tx *gorm.DB, productID int, numAdd int) error
}
