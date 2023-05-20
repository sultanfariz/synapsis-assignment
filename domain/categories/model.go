package categories

import (
	"context"
	"time"
)

type Category struct {
	Id        int
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CategoriesRepositoryInterface interface {
	GetAll(ctx context.Context) ([]*Category, error)
	GetById(ctx context.Context, id int) (*Category, error)
	GetByName(ctx context.Context, name string) (*Category, error)
	Insert(ctx context.Context, category *Category) (*Category, error)
}

type CategoriesUsecaseInterface interface {
	GetAll(ctx context.Context) ([]*Category, error)
	GetById(ctx context.Context, id int) (*Category, error)
	Insert(ctx context.Context, category *Category) (*Category, error)
}
