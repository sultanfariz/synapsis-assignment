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

	// // admin routes
	// admin := v1.Group("", middleware.JWTWithConfig(controllers.JWTMiddleware))
	// {
	// 	// gym endpoint
	// 	admin.GET("/gyms/count", controllers.GymController.CountAll, middlewares.OperationalAdmin())
	// 	admin.POST("/gyms", controllers.GymController.Create, middlewares.Superadmin())
	// 	admin.PUT("/gyms/:gymId", controllers.GymController.Update, middlewares.OperationalAdmin())
	// 	admin.DELETE("/gyms/:gymId", controllers.GymController.Delete, middlewares.Superadmin())

	// 	// class endpoint
	// 	admin.GET("/classes/count", controllers.ClassController.CountAll, middlewares.OperationalAdmin())
	// 	admin.POST("/gyms/:gymId/classes", controllers.ClassController.Create, middlewares.OperationalAdmin())
	// 	admin.PUT("/gyms/:gymId/classes/:classId", controllers.ClassController.Update, middlewares.OperationalAdmin())
	// 	admin.DELETE("/gyms/:gymId/classes/:classId", controllers.ClassController.Delete, middlewares.OperationalAdmin())

	// 	// operational admin endpoint

	// 	admin.PUT("/admin", controllers.OperationaladminsController.UpdatePassword, middlewares.OperationalAdmin())

	// 	//membership endpoint
	// 	admin.POST("/memberships", controllers.MembershipsController.Insert, middlewares.Superadmin())
	// 	admin.PUT("/memberships/:Id", controllers.MembershipsController.Update, middlewares.Superadmin())
	// 	admin.DELETE("/memberships/:Id", controllers.MembershipsController.Delete, middlewares.Superadmin())

	// 	//users endpoint
	// 	admin.GET("/users/count", controllers.ProfileController.CountAll, middlewares.OperationalAdmin())

	// 	// newsletter endpoint
	// 	admin.POST("/newsletters", controllers.NewslettersController.Create, middlewares.OperationalAdmin())
	// 	admin.PUT("/newsletters/:Id", controllers.NewslettersController.Update, middlewares.OperationalAdmin())
	// 	admin.DELETE("/newsletters/:Id", controllers.NewslettersController.Delete, middlewares.OperationalAdmin())

	// 	// session endpoint
	// 	admin.PUT("/sessions/:id", controllers.SessionsController.Update, middlewares.Superadmin())
	// 	admin.DELETE("/sessions/:id", controllers.SessionsController.Delete, middlewares.Superadmin())
	// 	admin.PUT("/schedules/:id", controllers.SchedulesController.Update, middlewares.Superadmin())
	// 	admin.DELETE("/schedules:/:id", controllers.SchedulesController.Delete, middlewares.Superadmin())

	// 	// admin endpoint
	// 	admin.PUT("/superadmin", controllers.SuperadminsController.UpdatePassword, middlewares.Superadmin())
	// 	admin.GET("/superadmin/admin/count", controllers.OperationaladminsController.CountAll, middlewares.Superadmin())
	// 	admin.GET("/superadmin/admin", controllers.OperationaladminsController.GetAll, middlewares.Superadmin())
	// 	admin.DELETE("/superadmin/logout", controllers.AuthController.SuperadminLogout, middlewares.Superadmin())

	// 	// booking endpoint
	// 	admin.GET("/bookings/gym/:gymID", controllers.BookingDetailsController.GetByGymID, middlewares.OperationalAdmin())
	// 	admin.GET("/bookings/count", controllers.BookingDetailsController.CountAll, middlewares.OperationalAdmin())
	// 	admin.PUT("/booking/:bookingID", controllers.BookingDetailsController.Update, middlewares.Superadmin())
	// }
}
