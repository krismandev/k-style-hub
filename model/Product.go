package model

type Product struct {
	ID                int                `gorm:"id;primaryKey"`
	Name              string             `gorm:"name"`
	Description       string             `gorm:"description"`
	Status            int                `gorm:"status"`
	CreatedAt         string             `gorm:"created_at"`
	UpdatedAt         string             `gorm:"updated_at"`
	Price             float64            `gorm:"price"`
	Stock             Stock              `gorm:"foreignKey:ProductID;references:ID"`
	StockTransactions []StockTransaction `gorm:"foreignKey:ProductID;references:ID"`
}
