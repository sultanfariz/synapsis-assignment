package payments

import (
	"context"
	"time"
)

type Payment struct {
	Id        int
	Method    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaymentsRepositoryInterface interface {
	GetAll(ctx context.Context) ([]*Payment, error)
	GetById(ctx context.Context, id int) (*Payment, error)
	GetByName(ctx context.Context, name string) (*Payment, error)
	Insert(ctx context.Context, method *Payment) (*Payment, error)
}

type PaymentsUsecaseInterface interface {
	GetAll(ctx context.Context) ([]*Payment, error)
	GetById(ctx context.Context, id int) (*Payment, error)
	Insert(ctx context.Context, method *Payment) (*Payment, error)
}
