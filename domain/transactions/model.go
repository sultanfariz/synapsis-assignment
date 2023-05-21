package transactions

import (
	"context"
	"time"
)

type CheckoutList struct {
	ProductId     int `validate:"required"`
	TransactionId int `validate:"required"`
	Quantity      int `validate:"required"`
}

type Transaction struct {
	Id        int
	UserId    int `validate:"required"`
	StatusId  int
	PaymentId int `validate:"required"`
	TotalCost int
	Products  []CheckoutList
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionsRepositoryInterface interface {
	GetByUser(ctx context.Context, userId int) ([]*Transaction, error)
	GetById(ctx context.Context, id int) (*Transaction, error)
	Insert(ctx context.Context, trx *Transaction) (*Transaction, error)
	UpdateStatus(ctx context.Context, id int, statusId int) (*Transaction, error)
}

type CheckoutListRepositoryInterface interface {
	GetByTransaction(ctx context.Context, transactionId int) ([]CheckoutList, error)
	// Insert(ctx context.Context, checkoutList *CheckoutList) (*CheckoutList, error)
}

type TransactionsUsecaseInterface interface {
	GetByUser(ctx context.Context, userId int) ([]*Transaction, error)
	GetById(ctx context.Context, id int) (*Transaction, error)
	Insert(ctx context.Context, trx *Transaction) (*Transaction, error)
	UpdateStatus(ctx context.Context, id int, statusId int, userId int) (*Transaction, error)
}
