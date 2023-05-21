package payments

import (
	"context"
	"strings"
	"time"

	model "github.com/sultanfariz/synapsis-assignment/domain/payments"

	"gorm.io/gorm"
)

type PaymentsRepository struct {
	DBConnection *gorm.DB
}

func NewPaymentsRepository(db *gorm.DB) *PaymentsRepository {
	return &PaymentsRepository{
		DBConnection: db,
	}
}

func (r *PaymentsRepository) GetAll(ctx context.Context) ([]*model.Payment, error) {
	data := []*model.Payment{}
	if err := r.DBConnection.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *PaymentsRepository) GetByName(ctx context.Context, method string) (*model.Payment, error) {
	data := model.Payment{}
	method = strings.ToLower(method)
	if err := r.DBConnection.Where("method = ?", method).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *PaymentsRepository) GetById(ctx context.Context, id int) (*model.Payment, error) {
	data := model.Payment{}
	if err := r.DBConnection.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *PaymentsRepository) Insert(ctx context.Context, payment *model.Payment) (*model.Payment, error) {
	payment.Method = strings.ToLower(payment.Method)
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()

	if err := r.DBConnection.Create(&payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}
