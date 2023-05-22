package categories

import (
	"context"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
)

type CategoriesUsecase struct {
	CategoriesRepository CategoriesRepositoryInterface
	ContextTimeout       time.Duration
}

func NewCategoriesUsecase(cr CategoriesRepositoryInterface, timeout time.Duration) *CategoriesUsecase {
	return &CategoriesUsecase{
		CategoriesRepository: cr,
		ContextTimeout:       timeout,
	}
}

func (cu *CategoriesUsecase) Insert(ctx context.Context, category Category) (*Category, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	validator := validator.New()
	if err := validator.Struct(category); err != nil {
		return nil, err
	}
	if category.Category == "" {
		return nil, commons.ErrEmptyInput
	}

	category.Category = strings.ToLower(category.Category)

	// check if category is already exist
	categoryRes, err := cu.CategoriesRepository.GetByName(ctx, category.Category)
	if err == nil && categoryRes != nil {
		return nil, commons.ErrCategoryAlreadyExists
	}

	// insert category to db
	categoryRes, err = cu.CategoriesRepository.Insert(ctx, &category)
	if err != nil {
		return nil, err
	}

	return categoryRes, nil
}

func (cu *CategoriesUsecase) GetAll(ctx context.Context) ([]*Category, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	// get all categories from db
	categories, err := cu.CategoriesRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cu *CategoriesUsecase) GetById(ctx context.Context, id int) (*Category, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	// get category by id from db
	category, err := cu.CategoriesRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
