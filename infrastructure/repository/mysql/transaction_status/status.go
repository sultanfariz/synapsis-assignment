package transaction_status

import (
	"context"
	"strings"
	"time"

	model "github.com/sultanfariz/synapsis-assignment/domain/transaction_status"

	"gorm.io/gorm"
)

type TransactionStatusRepository struct {
	DBConnection *gorm.DB
}

func NewTransactionStatusRepository(db *gorm.DB) *TransactionStatusRepository {
	return &TransactionStatusRepository{
		DBConnection: db,
	}
}

func (r *TransactionStatusRepository) GetAll(ctx context.Context) ([]*model.TransactionStatus, error) {
	data := []*model.TransactionStatus{}
	if err := r.DBConnection.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *TransactionStatusRepository) GetByName(ctx context.Context, status string) (*model.TransactionStatus, error) {
	data := model.TransactionStatus{}
	status = strings.ToLower(status)
	if err := r.DBConnection.Where("status = ?", status).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *TransactionStatusRepository) GetById(ctx context.Context, id int) (*model.TransactionStatus, error) {
	data := model.TransactionStatus{}
	if err := r.DBConnection.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *TransactionStatusRepository) Insert(ctx context.Context, trxStatus *model.TransactionStatus) (*model.TransactionStatus, error) {
	trxStatus.Status = strings.ToLower(trxStatus.Status)
	trxStatus.CreatedAt = time.Now()
	trxStatus.UpdatedAt = time.Now()

	if err := r.DBConnection.Create(&trxStatus).Error; err != nil {
		return nil, err
	}

	return trxStatus, nil
}
