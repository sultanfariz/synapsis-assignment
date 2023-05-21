package transaction_status

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
)

type TransactionStatusUsecase struct {
	TransactionStatusRepository TransactionStatusRepositoryInterface
	ContextTimeout              time.Duration
}

func NewTransactionStatusUsecase(tsr TransactionStatusRepositoryInterface, timeout time.Duration) *TransactionStatusUsecase {
	return &TransactionStatusUsecase{
		TransactionStatusRepository: tsr,
		ContextTimeout:              timeout,
	}
}

func (cu *TransactionStatusUsecase) Insert(ctx context.Context, trxStatus TransactionStatus) (*TransactionStatus, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	validator := validator.New()
	if err := validator.Struct(trxStatus); err != nil {
		return nil, err
	}
	if trxStatus.Status == "" {
		return nil, commons.ErrEmptyInput
	}

	// check if trxStatus method is already exist
	trxStatusRes, err := cu.TransactionStatusRepository.GetByName(ctx, trxStatus.Status)
	if err == nil && trxStatusRes != nil {
		return nil, commons.ErrTransactionStatusAlreadyExists
	}

	// insert trxStatus to db
	trxStatusRes, err = cu.TransactionStatusRepository.Insert(ctx, &trxStatus)
	if err != nil {
		return nil, err
	}

	return trxStatusRes, nil
}

func (cu *TransactionStatusUsecase) GetAll(ctx context.Context) ([]*TransactionStatus, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	// get all transaction_status from db
	transaction_status, err := cu.TransactionStatusRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return transaction_status, nil
}

func (cu *TransactionStatusUsecase) GetById(ctx context.Context, id int) (*TransactionStatus, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	// get trxStatus by id from db
	trxStatus, err := cu.TransactionStatusRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return trxStatus, nil
}
