package controller

import "github.com/labstack/echo/v4"

type OrderController interface {
	CreateOrder(c echo.Context) error
	CancelOrder(c echo.Context) error
	GetOrder(c echo.Context) error
	GetDetailOrder(c echo.Context) error
	UpdateOrder(c echo.Context) error
}
