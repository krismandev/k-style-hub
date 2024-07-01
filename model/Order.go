package model

type Order struct {
	ID            int           `gorm:"id"`
	Code          string        `gorm:"code"`
	GrandTotal    float64       `gorm:"grand_total"`
	CustomerID    int           `gorm:"customer_id"`
	AddressID     int           `gorm:"address_id"`
	PaymentStatus int           `gorm:"payment_status"`
	OrderStatus   int           `gorm:"order_status"`
	CreatedAt     *string       `gorm:"created_at"`
	UpdatedAt     *string       `gorm:"updated_at"`
	OrderDetail   []OrderDetail `gorm:"foreignKey:OrderID;references:ID"`
}
