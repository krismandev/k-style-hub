package app

import (
	"k-style-test/app/http/controller"
	"k-style-test/app/http/route"
	"k-style-test/config"
	"k-style-test/repository"
	"k-style-test/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Application struct {
	DB            *gorm.DB
	Validate      *validator.Validate
	Echo          *echo.Echo
	Config        config.Configuration
	WsConnections []*config.WebSocketConnection
}

func InitApp(app *Application) {
	userRepo := repository.NewUserRepository(app.DB)
	authLogRepo := repository.NewAuthRepository(app.DB)
	addressRepository := repository.NewAddressRepository(app.DB)
	customerRepository := repository.NewCustomerRepository(app.DB)
	customerUseCase := usecase.NewCustomerUseCase(customerRepository, userRepo, app.Validate, app.DB)
	customerController := controller.NewCustomerController(customerUseCase)

	authUseCase := usecase.NewAuthUseCase(app.Validate, userRepo, customerRepository, authLogRepo)
	authController := controller.NewAuthController(authUseCase)

	productRepository := repository.NewProductRepository(app.DB)
	stockRepository := repository.NewStockRepository(app.DB)

	orderRepository := repository.NewOrderRepository(app.DB)
	orderUseCase := usecase.NewOrderUseCase(app.DB, orderRepository, app.Validate, customerRepository, productRepository, stockRepository, addressRepository)
	orderController := controller.NewOrderController(orderUseCase)

	// chatroomRepository := repository.NewChatroomRepository(app.DB)
	// chatroomUseCase := usecase.NewChatroomUseCase(app.Validate, chatroomRepository, userRepository)
	// chatroomController := controller.NewChatroomController(chatroomUseCase)

	// chatRepository := repository.NewChatRepository(app.DB)
	// chatUseCase := usecase.NewChatUseCase(app.Validate, chatRepository, chatroomRepository, userRepository)

	// wsController := controller.NewWsController(userUseCase, &app.WsConnections, chatUseCase)

	routeConfig := route.RouteConfig{
		Echo:               app.Echo,
		CustomerController: customerController,
		AuthController:     authController,
		OrderController:    orderController,
	}

	routeConfig.Setup()

	// go wsController.HandleIncommingMessage()
}
