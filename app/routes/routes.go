package routes

import (
	"github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/auth"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/categories"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/products"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllersList struct {
	JWTMiddleware        middleware.JWTConfig
	AuthController       *auth.Controllers
	ProductsController   *products.Controllers
	CategoriesController *categories.Controllers
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
		// v1.GET("/products", controllers.ProductsController.GetById)

		// category endpoints
		v1.GET("/categories", controllers.CategoriesController.GetAll)
		v1.POST("/categories", controllers.CategoriesController.Insert)

		// // class endpoint
		// v1.GET("/classes", controllers.ClassController.GetAll)
		// v1.GET("/classes/:classId", controllers.ClassController.GetById)

		// auth endpoints
		v1.POST("/login", controllers.AuthController.Login)
		v1.POST("/register", controllers.AuthController.Register)
	}

	// Member routes
	// user := v1.Group("", middleware.JWTWithConfig(controllers.JWTMiddleware))
	// {
	// 	user.GET("/account/:id/mybookings", controllers.BookingDetailsController.GetByUserID, middlewares.Member())
	// 	user.GET("/bookings/:id", controllers.BookingDetailsController.GetByID, middlewares.Member())
	// 	user.GET("/account/:id", controllers.UsersController.GetByID, middlewares.Member())
	// 	user.PUT("/account", controllers.UsersController.Update, middlewares.Member())
	// 	user.PUT("/mybooking/:bookingID", controllers.BookingDetailsController.Update, middlewares.Member())
	// }

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
