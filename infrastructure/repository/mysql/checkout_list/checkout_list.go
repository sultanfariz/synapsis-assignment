package checkout_list

import (
	"context"

	model "github.com/sultanfariz/synapsis-assignment/domain/transactions"

	"gorm.io/gorm"
)

type CheckoutListRepository struct {
	DBConnection *gorm.DB
}

func NewCheckoutListRepository(db *gorm.DB) *CheckoutListRepository {
	return &CheckoutListRepository{
		DBConnection: db,
	}
}

func (r *CheckoutListRepository) GetByTransaction(ctx context.Context, transactionId int) ([]model.CheckoutList, error) {
	data := []model.CheckoutList{}
	if err := r.DBConnection.Where("transaction_id = ?", transactionId).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *CheckoutListRepository) GetById(ctx context.Context, id int) (*model.CheckoutList, error) {
	data := model.CheckoutList{}
	if err := r.DBConnection.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// func (r *CheckoutListRepository) Insert(ctx context.Context, checkoutList *model.CheckoutList) (*model.CheckoutList, error) {
// 	if err := r.DBConnection.Create(&checkoutList).Error; err != nil {
// 		return nil, err
// 	}

// 	return checkoutList, nil
// }
