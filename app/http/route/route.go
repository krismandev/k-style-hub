package route

import (
	"k-style-test/app/http/controller"
	"k-style-test/app/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	Echo               *echo.Echo
	AuthMiddleware     gin.HandlerFunc
	CustomerController controller.CustomerController
	AuthController     controller.AuthController
	OrderController    controller.OrderController
}

func (c *RouteConfig) Setup() {
	c.SetUpPrivateRoute()
	c.SetUpPublicRoute()
}

func (rc *RouteConfig) SetUpPrivateRoute() {
	orderRoute := rc.Echo.Group("order")
	{
		orderRoute.Use(middleware.JWTAuth())
		orderRoute.GET("", rc.OrderController.GetOrder)
		orderRoute.GET("/:id", rc.OrderController.GetDetailOrder)
		orderRoute.POST("", rc.OrderController.CreateOrder)
		orderRoute.DELETE("", rc.OrderController.CancelOrder)
		orderRoute.PATCH("", rc.OrderController.UpdateOrder)
	}

	customerRoute := rc.Echo.Group("customer")
	{
		customerRoute.Use(middleware.JWTAuth())
		customerRoute.GET("", rc.CustomerController.GetCustomer)
		customerRoute.PATCH("", rc.CustomerController.UpdateCustomer)
		customerRoute.DELETE("", rc.CustomerController.DeleteCustomer)
	}

	rc.Echo.GET("/auth-log", rc.AuthController.GetAuthLog)
}

func (rc *RouteConfig) SetUpPublicRoute() {
	publicRoute := rc.Echo.Group("")
	publicRoute.POST("/register", rc.CustomerController.Register)

	publicRoute.POST("/login", rc.AuthController.Login)
}
