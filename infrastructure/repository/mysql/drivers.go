package mysql

import (
	"fmt"

	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/carts"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/categories"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/products"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_USERNAME": "root",
		"DB_PASSWORD": "",
		"DB_HOST":     "localhost",
		"DB_PORT":     "3306",
		"DB_NAME":     "synapsis_assignment",
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config["DB_USERNAME"], config["DB_PASSWORD"], config["DB_HOST"], config["DB_PORT"], config["DB_NAME"])

	DB, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&users.User{},
		&products.Product{},
		&categories.Category{},
		&carts.Cart{},
	)
	return DB
}
