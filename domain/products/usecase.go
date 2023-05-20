package products

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sultanfariz/synapsis-assignment/domain/categories"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
)

type ProductsUsecase struct {
	ProductsRepository   ProductsRepositoryInterface
	CategoriesRepository categories.CategoriesRepositoryInterface
	ContextTimeout       time.Duration
}

func NewProductsUsecase(pr ProductsRepositoryInterface, cr categories.CategoriesRepositoryInterface, timeout time.Duration) *ProductsUsecase {
	return &ProductsUsecase{
		ProductsRepository:   pr,
		CategoriesRepository: cr,
		ContextTimeout:       timeout,
	}
}

func (pu *ProductsUsecase) GetByCategory(ctx context.Context, name string) ([]*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()

	if name == "" {
		return nil, commons.ErrEmptyInput
	}

	// get category id from db
	category, err := pu.CategoriesRepository.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	// get all products by category from db
	products, err := pu.ProductsRepository.GetByCategory(ctx, category.Id)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (pu *ProductsUsecase) Insert(ctx context.Context, product *Product) (*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()

	validator := validator.New()
	if err := validator.Struct(product); err != nil {
		return nil, err
	}

	// insert product to db
	product, err := pu.ProductsRepository.Insert(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductsUsecase) GetAll(ctx context.Context) ([]*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()

	// get all categories from db
	categories, err := pu.ProductsRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (pu *ProductsUsecase) GetById(ctx context.Context, id int) (*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.ContextTimeout)
	defer cancel()

	// get product by id from db
	product, err := pu.ProductsRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
