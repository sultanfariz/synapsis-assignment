package main

import (
	"log"
	"time"

	"github.com/sultanfariz/synapsis-assignment/app/routes"
	_cartsUsecase "github.com/sultanfariz/synapsis-assignment/domain/carts"
	_categoriesUsecase "github.com/sultanfariz/synapsis-assignment/domain/categories"
	_productsUsecase "github.com/sultanfariz/synapsis-assignment/domain/products"
	_usersUsecase "github.com/sultanfariz/synapsis-assignment/domain/users"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
	_cartsRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/carts"
	_categoriesRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/categories"
	_productsRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/products"
	_usersRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/users"
	_authController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/auth"
	_cartsController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/carts"
	_categoriesController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/categories"
	_productsController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/products"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql"
)

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if viper.GetBool("debug") {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := mysql.InitDB()
	configJWT := commons.ConfigJWT{
		SecretJWT:       viper.GetString("JWT_SECRET"),
		ExpiresDuration: viper.GetInt("JWT_EXPIRED"),
	}

	timeoutContext := time.Duration(viper.GetInt("SERVER_TIMEOUT")) * time.Second

	// Users initialize
	usersRepo := _usersRepo.NewUsersRepository(db)
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepo, timeoutContext, &configJWT)

	// Categories initialize
	categoriesRepo := _categoriesRepo.NewCategoriesRepository(db)
	categoriesUsecase := _categoriesUsecase.NewCategoriesUsecase(categoriesRepo, timeoutContext)

	// Products initialize
	productsRepo := _productsRepo.NewProductsRepository(db)
	productsUsecase := _productsUsecase.NewProductsUsecase(productsRepo, categoriesRepo, timeoutContext)

	// Carts initialize
	cartsRepo := _cartsRepo.NewCartsRepository(db)
	cartsUsecase := _cartsUsecase.NewCartsUsecase(cartsRepo, productsRepo, timeoutContext)

	// Auth initialize
	authController := _authController.NewControllers(*usersUsecase)
	productsController := _productsController.NewControllers(*productsUsecase)
	categoriesController := _categoriesController.NewControllers(*categoriesUsecase)
	cartsController := _cartsController.NewControllers(*cartsUsecase)

	routesInit := routes.ControllersList{
		JWTMiddleware:        configJWT.Init(),
		AuthController:       authController,
		ProductsController:   productsController,
		CategoriesController: categoriesController,
		CartsController:      cartsController,
	}
	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("SERVER_PORT")))
}
