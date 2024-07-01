package response

import "k-style-test/model"

type OrderResponse struct {
	OrderID       int                    `json:"order_id"`
	Code          string                 `json:"code"`
	CustomerID    int                    `json:"customer_id"`
	GrandTotal    float64                `json:"grand_total"`
	AddressID     int                    `json:"address_id"`
	PaymentStatus int                    `json:"payment_status"`
	CreatedAt     string                 `json:"created_at"`
	OrderStatus   int                    `json:"order_status"`
	OrderDetails  *[]OrderDetailResponse `json:"order_details,omitempty"`
}

type CreateOrderResponse struct {
	OrderResponse
}

type UpdateOrderResponse struct {
	OrderResponse
}

type CancelOrderRespoonse struct {
	OrderResponse
}

type OrderDetailResponse struct {
	OrderDetailID int     `json:"order_detail_id"`
	ProductID     int     `json:"product_id"`
	OrderID       int     `json:"order_id"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
	Amount        float64 `json:"amount"`
	CreatedAt     string  `json:"created_at"`
	Status        int     `json:"status"`
}

func ToOrderResponse(order *model.Order) OrderResponse {
	var resp OrderResponse
	resp.Code = order.Code
	resp.OrderID = order.ID
	resp.AddressID = order.AddressID
	resp.GrandTotal = order.GrandTotal
	resp.OrderStatus = order.OrderStatus
	resp.PaymentStatus = order.PaymentStatus
	resp.CustomerID = order.CustomerID

	return resp
}

func ToOrderDetailResponse(order *model.Order) OrderResponse {
	var resp OrderResponse
	resp.Code = order.Code
	resp.OrderID = order.ID
	resp.AddressID = order.AddressID
	resp.GrandTotal = order.GrandTotal
	resp.OrderStatus = order.OrderStatus
	resp.PaymentStatus = order.PaymentStatus
	resp.CustomerID = order.CustomerID

	var details []OrderDetailResponse
	for _, each := range order.OrderDetail {
		dt := OrderDetailResponse{}
		dt.Amount = each.Amount
		dt.Price = each.Price
		dt.CreatedAt = each.CreatedAt
		dt.OrderID = each.OrderID
		dt.Quantity = each.Quantity
		dt.Status = each.Status
		dt.ProductID = each.ProductID
		details = append(details, dt)
	}
	resp.OrderDetails = &details

	return resp
}

func ToCancelOrderResponse(order *model.Order) CancelOrderRespoonse {
	var resp CancelOrderRespoonse

	resp.OrderID = order.ID
	resp.Code = order.Code
	resp.CustomerID = order.CustomerID
	resp.GrandTotal = order.GrandTotal
	resp.AddressID = order.AddressID
	resp.CreatedAt = *order.CreatedAt
	resp.OrderStatus = order.OrderStatus
	resp.PaymentStatus = order.PaymentStatus
	return resp
}

func ToUpdateOrderResponse(order *model.Order) UpdateOrderResponse {
	var resp UpdateOrderResponse

	resp.OrderID = order.ID
	resp.Code = order.Code
	resp.CustomerID = order.CustomerID
	resp.GrandTotal = order.GrandTotal
	resp.AddressID = order.AddressID
	resp.CreatedAt = *order.CreatedAt
	resp.OrderStatus = order.OrderStatus
	resp.PaymentStatus = order.PaymentStatus
	return resp
}

func ToCreateOrderResponse(order *model.Order) CreateOrderResponse {
	var resp CreateOrderResponse

	resp.OrderID = order.ID
	resp.Code = order.Code
	resp.CustomerID = order.CustomerID
	resp.GrandTotal = order.GrandTotal
	resp.AddressID = order.AddressID
	resp.CreatedAt = *order.CreatedAt
	resp.OrderStatus = order.OrderStatus
	resp.PaymentStatus = order.PaymentStatus
	return resp
}
