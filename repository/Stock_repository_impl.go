package repository

import (
	"k-style-test/model"
	"k-style-test/utility"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StockRepositoryImpl struct {
	DB *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &StockRepositoryImpl{DB: db}
}

func (r *StockRepositoryImpl) GetProductStockById(stock *model.Stock) error {
	var err error
	result := r.DB.Where("product_id = ?", stock.ProductID).First(stock)
	err = utility.CheckErrorResult(result)
	if err != nil {

		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (r *StockRepositoryImpl) UpdateProductStock(tx *gorm.DB, productID int, stock int) error {
	var err error
	result := tx.Model(&model.Stock{}).Where("product_id = ?", productID).Update("quantity", stock)
	err = utility.CheckErrorResult(result)
	if err != nil {

		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (r *StockRepositoryImpl) CreateStockTransaction(tx *gorm.DB, stockTransaction *[]model.StockTransaction) error {
	var err error
	result := tx.Create(stockTransaction)
	err = utility.CheckErrorResult(result)
	if err != nil {

		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (r *StockRepositoryImpl) GetProductStockByProductIdListTransaction(tx *gorm.DB, stocks *[]model.Stock, productIds []int) error {
	var err error

	// pesimistic locking to avoid other request that came at the same time get same stock
	result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&stocks, productIds)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (r *StockRepositoryImpl) AddStock(tx *gorm.DB, productID int, numAdd int) error {
	var err error

	result := tx.Exec("UPDATE stocks SET quantity = quantity + ?, updated_at = now() WHERE product_id = ?", numAdd, productID)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}
