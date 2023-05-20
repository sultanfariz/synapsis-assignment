package main

import (
	"log"
	"time"

	"github.com/sultanfariz/synapsis-assignment/app/routes"
	_productsUsecase "github.com/sultanfariz/synapsis-assignment/domain/products"
	_usersUsecase "github.com/sultanfariz/synapsis-assignment/domain/users"
	_categoriesRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/categories"
	_productsRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/products"
	_usersRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/users"
	_authController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/auth"
	_productsController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/products"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
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
	// categoriesUsecase := _categoriesUsecase.NewCategoriesUsecase(categoriesRepo, timeoutContext)

	// Products initialize
	productsRepo := _productsRepo.NewProductsRepository(db)
	productsUsecase := _productsUsecase.NewProductsUsecase(productsRepo, categoriesRepo, timeoutContext)

	// Auth initialize
	authController := _authController.NewControllers(*usersUsecase)
	productsController := _productsController.NewControllers(*productsUsecase)

	routesInit := routes.ControllersList{
		JWTMiddleware:      configJWT.Init(),
		AuthController:     authController,
		ProductsController: productsController,
	}
	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("SERVER_PORT")))
}
