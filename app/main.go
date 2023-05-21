package main

import (
	"log"
	"time"

	"github.com/sultanfariz/synapsis-assignment/app/routes"
	_cartsUsecase "github.com/sultanfariz/synapsis-assignment/domain/carts"
	_categoriesUsecase "github.com/sultanfariz/synapsis-assignment/domain/categories"
	_paymentsUsecase "github.com/sultanfariz/synapsis-assignment/domain/payments"
	_productsUsecase "github.com/sultanfariz/synapsis-assignment/domain/products"
	_transactionStatusUsecase "github.com/sultanfariz/synapsis-assignment/domain/transaction_status"
	_transactionsUsecase "github.com/sultanfariz/synapsis-assignment/domain/transactions"
	_usersUsecase "github.com/sultanfariz/synapsis-assignment/domain/users"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
	_cartsRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/carts"
	_categoriesRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/categories"
	_checkoutListRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/checkout_list"
	_paymentsRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/payments"
	_productsRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/products"
	_transactionStatusRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/transaction_status"
	_transactionsRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/transactions"
	_usersRepo "github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/users"
	_authController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/auth"
	_cartsController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/carts"
	_categoriesController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/categories"
	_paymentsController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/payments"
	_productsController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/products"
	_transactionStatusController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/transaction_status"
	_transactionsController "github.com/sultanfariz/synapsis-assignment/infrastructure/transport/http/transactions"

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

	// Payments initialize
	paymentsRepo := _paymentsRepo.NewPaymentsRepository(db)
	paymentsUsecase := _paymentsUsecase.NewPaymentsUsecase(paymentsRepo, timeoutContext)

	// Transaction Status initialize
	transactionStatusRepo := _transactionStatusRepo.NewTransactionStatusRepository(db)
	transactionStatusUsecase := _transactionStatusUsecase.NewTransactionStatusUsecase(transactionStatusRepo, timeoutContext)

	// Transactions initialize
	transactionsRepo := _transactionsRepo.NewTransactionsRepository(db)
	checkoutListRepo := _checkoutListRepo.NewCheckoutListRepository(db)
	transactionsUsecase := _transactionsUsecase.NewTransactionsUsecase(transactionsRepo, checkoutListRepo, productsRepo, timeoutContext)

	// Controllers initialize
	authController := _authController.NewControllers(*usersUsecase)
	productsController := _productsController.NewControllers(*productsUsecase)
	categoriesController := _categoriesController.NewControllers(*categoriesUsecase)
	cartsController := _cartsController.NewControllers(*cartsUsecase)
	paymentsController := _paymentsController.NewControllers(*paymentsUsecase)
	transactionStatusController := _transactionStatusController.NewControllers(*transactionStatusUsecase)
	transactionsController := _transactionsController.NewControllers(*transactionsUsecase)

	routesInit := routes.ControllersList{
		JWTMiddleware:               configJWT.Init(),
		AuthController:              authController,
		ProductsController:          productsController,
		CategoriesController:        categoriesController,
		CartsController:             cartsController,
		PaymentsController:          paymentsController,
		TransactionStatusController: transactionStatusController,
		TransactionsController:      transactionsController,
	}
	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(viper.GetString("SERVER_PORT")))
}
