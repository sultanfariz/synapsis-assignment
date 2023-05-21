package carts

import (
	"time"

	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/products"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/users"
)

type Cart struct {
	Id        int              `gorm:"primaryKey"`
	UserId    int              `gorm:"column:user_id;not null"`
	User      users.User       `gorm:"foreignKey:UserId"`
	ProductId int              `gorm:"column:product_id;not null"`
	Product   products.Product `gorm:"foreignKey:ProductId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
