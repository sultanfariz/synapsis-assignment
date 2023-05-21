package products

import (
	"context"
	"time"

	model "github.com/sultanfariz/synapsis-assignment/domain/products"

	"gorm.io/gorm"
)

type ProductsRepository struct {
	DBConnection *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{
		DBConnection: db,
	}
}

func (r *ProductsRepository) GetAll(ctx context.Context) ([]*model.Product, error) {
	data := []*model.Product{}
	if err := r.DBConnection.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *ProductsRepository) GetByCategory(ctx context.Context, categoryId int) ([]*model.Product, error) {
	data := []*model.Product{}
	if err := r.DBConnection.Where("category_id = ?", categoryId).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *ProductsRepository) GetById(ctx context.Context, id int) (*model.Product, error) {
	data := model.Product{}
	if err := r.DBConnection.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ProductsRepository) GetByIds(ctx context.Context, ids []int) ([]*model.Product, error) {
	data := []*model.Product{}
	if err := r.DBConnection.Where("id IN ?", ids).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *ProductsRepository) Insert(ctx context.Context, user *model.Product) (*model.Product, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := r.DBConnection.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
