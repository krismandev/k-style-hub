package model

type OrderDetail struct {
	ID        int     `gorm:"id"`
	OrderID   int     `gorm:"order_id"`
	ProductID int     `gorm:"product_id"`
	Price     float64 `gorm:"price"`
	Quantity  int     `gorm:"qty"`
	Amount    float64 `gorm:"amount"`
	CreatedAt string  `gorm:"created_at"`
	Status    int     `gorm:"status"`
}
