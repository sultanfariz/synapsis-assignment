package transaction_status

import (
	"context"
	"time"
)

type TransactionStatus struct {
	Id        int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionStatusRepositoryInterface interface {
	GetAll(ctx context.Context) ([]*TransactionStatus, error)
	GetById(ctx context.Context, id int) (*TransactionStatus, error)
	GetByName(ctx context.Context, name string) (*TransactionStatus, error)
	Insert(ctx context.Context, status *TransactionStatus) (*TransactionStatus, error)
}

type TransactionStatusUsecaseInterface interface {
	GetAll(ctx context.Context) ([]*TransactionStatus, error)
	GetById(ctx context.Context, id int) (*TransactionStatus, error)
	Insert(ctx context.Context, status *TransactionStatus) (*TransactionStatus, error)
}
