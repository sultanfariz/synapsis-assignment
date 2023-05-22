package carts

import (
	"context"
	"time"

	"github.com/sultanfariz/synapsis-assignment/domain/products"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
)

type CartsUsecase struct {
	CartsRepository    CartsRepositoryInterface
	ProductsRepository products.ProductsRepositoryInterface
	ContextTimeout     time.Duration
}

func NewCartsUsecase(cr CartsRepositoryInterface, pr products.ProductsRepositoryInterface, timeout time.Duration) *CartsUsecase {
	return &CartsUsecase{
		CartsRepository:    cr,
		ProductsRepository: pr,
		ContextTimeout:     timeout,
	}
}

func (cu *CartsUsecase) Insert(ctx context.Context, userId int, productId int) (*Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	if userId == 0 || productId == 0 {
		return nil, commons.ErrEmptyInput
	}

	// insert product cart to db
	product, err := cu.CartsRepository.Insert(ctx, userId, productId)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (cu *CartsUsecase) GetById(ctx context.Context, id int) (*Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	// get product by id from db
	product, err := cu.CartsRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (cu *CartsUsecase) GetByUser(ctx context.Context, userId int) ([]*products.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	// get product carts by user from db
	carts, err := cu.CartsRepository.GetByUser(ctx, userId)
	if err != nil {
		// check if record not found
		if err.Error() == "record not found" {
			return nil, commons.ErrCartIsEmpty
		}
		return nil, err
	}

	// map product id to array
	productIds := make([]int, 0)
	for _, cart := range carts {
		productIds = append(productIds, cart.ProductId)
	}

	products, err := cu.ProductsRepository.GetByIds(ctx, productIds)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (cu *CartsUsecase) Delete(ctx context.Context, id int, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	// get product cart by id from db
	cart, err := cu.CartsRepository.GetById(ctx, id)
	if err != nil {
		return commons.ErrProductNotFound
	}

	// validate the user owner
	if cart.UserId != userId {
		return commons.ErrUnauthorized
	}

	// delete product cart by id from db
	err = cu.CartsRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
