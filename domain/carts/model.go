package carts

import (
	"context"
	"time"

	"github.com/sultanfariz/synapsis-assignment/domain/products"
)

type Cart struct {
	Id        int
	UserId    int
	ProductId int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartsRepositoryInterface interface {
	GetByUser(ctx context.Context, userId int) ([]*Cart, error)
	GetById(ctx context.Context, id int) (*Cart, error)
	Insert(ctx context.Context, userId int, productId int) (*Cart, error)
}

type CartsUsecaseInterface interface {
	GetByUser(ctx context.Context, userId int) ([]*products.Product, error)
	GetById(ctx context.Context, id int) (*Cart, error)
	Insert(ctx context.Context, userId int, productId int) (*Cart, error)
}
