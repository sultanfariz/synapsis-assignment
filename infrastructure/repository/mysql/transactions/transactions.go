package transactions

import (
	"context"
	"time"

	"github.com/sultanfariz/synapsis-assignment/domain/products"
	model "github.com/sultanfariz/synapsis-assignment/domain/transactions"

	"gorm.io/gorm"
)

type TransactionsRepository struct {
	DBConnection *gorm.DB
}

func NewTransactionsRepository(db *gorm.DB) *TransactionsRepository {
	return &TransactionsRepository{
		DBConnection: db,
	}
}

func (r *TransactionsRepository) GetByUser(ctx context.Context, userId int) ([]*model.Transaction, error) {
	data := []*model.Transaction{}
	if err := r.DBConnection.Where("user_id = ?", userId).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *TransactionsRepository) GetById(ctx context.Context, id int) (*model.Transaction, error) {
	data := model.Transaction{}
	if err := r.DBConnection.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *TransactionsRepository) Insert(ctx context.Context, trx *model.Transaction) (*model.Transaction, error) {
	trx.CreatedAt = time.Now()
	trx.UpdatedAt = time.Now()

	tx := r.DBConnection.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&trx).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// subtract product stock
	for _, product := range trx.Products {
		// find product by id
		productObj := products.Product{}
		if err := tx.Where("id = ?", product.ProductId).First(&productObj).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// update product stock
		if err := tx.Model(&productObj).Where("id = ?", product.ProductId).Update("stock", gorm.Expr("stock - ?", product.Quantity)).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()

	return trx, nil
}

func (r *TransactionsRepository) UpdateStatus(ctx context.Context, id int, statusId int) (*model.Transaction, error) {
	trx := model.Transaction{}
	if err := r.DBConnection.Where("id = ?", id).First(&trx).Error; err != nil {
		return nil, err
	}

	trx.StatusId = statusId
	trx.UpdatedAt = time.Now()

	if err := r.DBConnection.Save(&trx).Error; err != nil {
		return nil, err
	}

	return &trx, nil
}
