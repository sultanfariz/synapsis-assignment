package carts

import (
	"context"
	"time"

	model "github.com/sultanfariz/synapsis-assignment/domain/carts"

	"gorm.io/gorm"
)

type CartsRepository struct {
	DBConnection *gorm.DB
}

func NewCartsRepository(db *gorm.DB) *CartsRepository {
	return &CartsRepository{
		DBConnection: db,
	}
}

func (r *CartsRepository) GetByUser(ctx context.Context, userId int) ([]*model.Cart, error) {
	data := []*model.Cart{}
	if err := r.DBConnection.Where("user_id = ?", userId).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *CartsRepository) GetById(ctx context.Context, id int) (*model.Cart, error) {
	data := model.Cart{}
	if err := r.DBConnection.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *CartsRepository) Insert(ctx context.Context, userId int, productId int) (*model.Cart, error) {
	cart := model.Cart{
		UserId:    userId,
		ProductId: productId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := r.DBConnection.Create(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *CartsRepository) Delete(ctx context.Context, id int) error {
	if err := r.DBConnection.Where("id = ?", id).Delete(&model.Cart{}).Error; err != nil {
		return err
	}

	return nil
}
