package request

type CreateOrderRequest struct {
	AddressID    int                  `json:"address_id" validate:"required"`
	UserID       int                  `json:"-"`
	CustomerID   int                  `json:"-"`
	OrderDetails []OrderDetailRequest `json:"order_details" validate:"required,min=1,dive,required"`
}

type UpdateOrderRequest struct {
	AddressID    int                  `json:"address_id" validate:"required"`
	OrderID      int                  `json:"order_id" validate:"required"`
	UserID       int                  `json:"-"`
	CustomerID   int                  `json:"-"`
	OrderDetails []OrderDetailRequest `json:"order_details" validate:"required,min=1,dive,required"`
}

type OrderDetailRequest struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required,min=1"`
}

type CancelOrderRequest struct {
	OrderID    int `json:"order_id" validate:"required"`
	CustomerID int `json:"-"`
}

type GetOrderRequest struct {
	OrderID    int            `json:"order_id"`
	CustomerID int            `json:"-"`
	Code       string         `json:"code"`
	Param      DataTableParam `json:"param"`
}
