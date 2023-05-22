package categories

import (
	"context"
	"strings"
	"time"

	model "github.com/sultanfariz/synapsis-assignment/domain/categories"

	"gorm.io/gorm"
)

type CategoriesRepository struct {
	DBConnection *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) *CategoriesRepository {
	return &CategoriesRepository{
		DBConnection: db,
	}
}

func (r *CategoriesRepository) GetAll(ctx context.Context) ([]*model.Category, error) {
	data := []*model.Category{}
	if err := r.DBConnection.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *CategoriesRepository) GetByName(ctx context.Context, name string) (*model.Category, error) {
	data := model.Category{}
	name = strings.ToLower(name)
	if err := r.DBConnection.Where("category = ?", name).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *CategoriesRepository) GetById(ctx context.Context, id int) (*model.Category, error) {
	data := model.Category{}
	if err := r.DBConnection.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *CategoriesRepository) Insert(ctx context.Context, category *model.Category) (*model.Category, error) {
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	if err := r.DBConnection.Create(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
