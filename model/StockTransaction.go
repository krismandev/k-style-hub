package model

type StockTransaction struct {
	ID              int    `json:"id"`
	ProductID       int    `gorm:"product_id"`
	QuantityChange  int    `gorm:"quantity_change"`
	TransactionType int    `gorm:"transaction_type"`
	ReferenceID     string `gorm:"reference_id"`
	LastStock       int    `gorm:"last_stock"`
	StockAfter      int    `gorm:"stock_after"`
	CreatedAt       string `gorm:"created_at"`
}
