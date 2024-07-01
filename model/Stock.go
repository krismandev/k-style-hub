package model

type Stock struct {
	ID        int    `gorm:"id;primaryKey"`
	ProductID int    `gorm:"product_id"`
	Quantity  int    `gorm:"quantity"`
	CreatedAt string `gorm:"created_at"`
	UpdatedAt string `gorm:"updated_at"`
}
