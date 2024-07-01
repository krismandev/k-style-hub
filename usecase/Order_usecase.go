package usecase

import (
	"context"
	"k-style-test/model/request"
	"k-style-test/model/response"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, request request.CreateOrderRequest) (response.CreateOrderResponse, error)
	CancelOrder(ctx context.Context, request request.CancelOrderRequest) (response.CancelOrderRespoonse, error)
	GetOrder(ctx context.Context, request request.GetOrderRequest) (lists []response.OrderResponse, totalData int64, err error)
	GetDetailOrder(ctx context.Context, request request.GetOrderRequest) (response.OrderResponse, error)
	UpdateOrder(ctx context.Context, request request.UpdateOrderRequest) (response.UpdateOrderResponse, error)
}
