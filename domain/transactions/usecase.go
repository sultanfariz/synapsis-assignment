package transactions

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sultanfariz/synapsis-assignment/domain/products"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
)

type TransactionsUsecase struct {
	TransactionsRepository TransactionsRepositoryInterface
	ProductsRepository     products.ProductsRepositoryInterface
	CheckoutListRepository CheckoutListRepositoryInterface
	ContextTimeout         time.Duration
}

func NewTransactionsUsecase(tr TransactionsRepositoryInterface, clr CheckoutListRepositoryInterface, pr products.ProductsRepositoryInterface, timeout time.Duration) *TransactionsUsecase {
	return &TransactionsUsecase{
		TransactionsRepository: tr,
		ProductsRepository:     pr,
		CheckoutListRepository: clr,
		ContextTimeout:         timeout,
	}
}

func (tu *TransactionsUsecase) Insert(ctx context.Context, trx Transaction) (*Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.ContextTimeout)
	defer cancel()

	validator := validator.New()
	if err := validator.Struct(trx); err != nil {
		return nil, err
	}

	// calculate total cost
	totalCost := 0
	for _, productEl := range trx.Products {
		// get product by id from db
		product, err := tu.ProductsRepository.GetById(ctx, productEl.ProductId)
		if err != nil {
			if err.Error() == "record not found" {
				return nil, commons.ErrProductNotFound
			}
			return nil, err
		}

		if productEl.Quantity > product.Stock {
			return nil, commons.ErrProductOutOfStock
		}

		totalCost += (product.Price * productEl.Quantity)
	}
	trx.TotalCost = totalCost
	trx.StatusId = 1

	// insert new transaction to db
	trxObj, err := tu.TransactionsRepository.Insert(ctx, &trx)
	if err != nil {
		return nil, err
	}

	return trxObj, nil
}

func (tu *TransactionsUsecase) GetById(ctx context.Context, id int) (*Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.ContextTimeout)
	defer cancel()

	// get trx by id from db
	trx, err := tu.TransactionsRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	// get product list by trx id from db
	productList, err := tu.CheckoutListRepository.GetByTransaction(ctx, id)
	if err != nil {
		return nil, err
	}
	trx.Products = productList

	return trx, nil
}

func (tu *TransactionsUsecase) GetByUser(ctx context.Context, userId int) ([]*Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.ContextTimeout)
	defer cancel()

	// get trx by user id from db
	trx, err := tu.TransactionsRepository.GetByUser(ctx, userId)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, commons.ErrTransactionNotFound
		}
		return nil, err
	}

	// get product list by trx id from db
	for _, trxEl := range trx {
		productList, err := tu.CheckoutListRepository.GetByTransaction(ctx, trxEl.Id)
		if err != nil {
			if err.Error() == "record not found" {
				return nil, commons.ErrTransactionNotFound
			}
			return nil, err
		}
		trxEl.Products = productList
	}

	return trx, nil
}

func (tu *TransactionsUsecase) UpdateStatus(ctx context.Context, id int, statusId int, userId int) (*Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.ContextTimeout)
	defer cancel()

	// update trx status
	trx, err := tu.TransactionsRepository.UpdateStatus(ctx, id, statusId)
	if err != nil {
		return nil, err
	}

	return trx, nil
}
