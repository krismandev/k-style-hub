package repository

import (
	"k-style-test/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrder(orders *[]model.Order, params map[string]interface{}) error
	GetOrderByID(order *model.Order) error
	CreateOrder(tx *gorm.DB, order *model.Order) error
	CreateOrderDetail(tx *gorm.DB, orderDetails *[]model.OrderDetail) error
	UpdateGrandTotal(tx *gorm.DB, orderID int, grandTotal float64) error
	UpdateOrder(tx *gorm.DB, order *model.Order) error
	UpdateOrderDetail(tx *gorm.DB, orderDetail *model.OrderDetail) error
	CancelOrder(tx *gorm.DB, order *model.Order) error
	OrderCount(customerID int, result *int64) error
}
