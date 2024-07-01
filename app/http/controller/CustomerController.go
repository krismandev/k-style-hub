package controller

import (
	"github.com/labstack/echo/v4"
)

type CustomerController interface {
	Register(c echo.Context) error
	// GetUsers(c echo.Context) error
	GetCustomer(c echo.Context) error
	UpdateCustomer(c echo.Context) error
	DeleteCustomer(c echo.Context) error
}
