package routes

import (
	"github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/auth"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/carts"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/categories"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/payments"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/products"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/transaction_status"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/transactions"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllersList struct {
	JWTMiddleware               echojwt.Config
	AuthController              *auth.Controllers
	ProductsController          *products.Controllers
	CategoriesController        *categories.Controllers
	CartsController             *carts.Controllers
	PaymentsController          *payments.Controllers
	TransactionStatusController *transaction_status.Controllers
	TransactionsController      *transactions.Controllers
}

func (controllers ControllersList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	v1 := e.Group("/api/v1")
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

	// unprotected routes
	{
		// product endpoints
		v1.GET("/products", controllers.ProductsController.GetAll)
		v1.POST("/products", controllers.ProductsController.Insert)

		// category endpoints
		v1.GET("/categories", controllers.CategoriesController.GetAll)
		v1.POST("/categories", controllers.CategoriesController.Insert)

		// payment endpoints
		v1.GET("/payments", controllers.PaymentsController.GetAll)
		v1.POST("/payments", controllers.PaymentsController.Insert)

		// trx status endpoints
		v1.GET("/trx-status", controllers.TransactionStatusController.GetAll)
		v1.POST("/trx-status", controllers.TransactionStatusController.Insert)

		// auth endpoints
		v1.POST("/login", controllers.AuthController.Login)
		v1.POST("/register", controllers.AuthController.Register)
	}

	// Users routes
	user := v1.Group("", echojwt.WithConfig(controllers.JWTMiddleware))
	{
		// carts endpoints
		user.POST("/carts", controllers.CartsController.Insert)
		user.GET("/carts", controllers.CartsController.GetByUser)
		user.DELETE("/carts/:id", controllers.CartsController.Delete)

		// transaction endpoints
		user.POST("/transactions", controllers.TransactionsController.Insert)
		user.GET("/transactions", controllers.TransactionsController.GetByUser)
		user.PUT("/transactions/:id", controllers.TransactionsController.UpdateStatus)
	}
}
