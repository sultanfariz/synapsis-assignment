package products

import (
	"context"
	"time"
)

type Product struct {
	Id          int
	Name        string `validate:"required"`
	Price       int    `validate:"required"`
	Stock       int    `validate:"required"`
	Description string `validate:"required"`
	PictureUrl  string
	CategoryId  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductsRepositoryInterface interface {
	GetByCategory(ctx context.Context, categoryId int) ([]*Product, error)
	GetAll(ctx context.Context) ([]*Product, error)
	GetById(ctx context.Context, id int) (*Product, error)
	GetByIds(ctx context.Context, ids []int) ([]*Product, error)
	Insert(ctx context.Context, category *Product) (*Product, error)
}

type ProductsUsecaseInterface interface {
	GetByCategory(ctx context.Context, category string) ([]*Product, error)
	GetAll(ctx context.Context) ([]*Product, error)
	GetById(ctx context.Context, id int) (*Product, error)
	Insert(ctx context.Context, category *Product) (*Product, error)
}
