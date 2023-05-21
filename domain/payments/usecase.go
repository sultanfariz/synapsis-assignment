package payments

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/commons"
)

type PaymentsUsecase struct {
	PaymentsRepository PaymentsRepositoryInterface
	ContextTimeout     time.Duration
}

func NewPaymentsUsecase(cr PaymentsRepositoryInterface, timeout time.Duration) *PaymentsUsecase {
	return &PaymentsUsecase{
		PaymentsRepository: cr,
		ContextTimeout:     timeout,
	}
}

func (cu *PaymentsUsecase) Insert(ctx context.Context, payment Payment) (*Payment, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	validator := validator.New()
	if err := validator.Struct(payment); err != nil {
		return nil, err
	}
	if payment.Method == "" {
		return nil, commons.ErrEmptyInput
	}

	// check if payment method is already exist
	paymentRes, err := cu.PaymentsRepository.GetByName(ctx, payment.Method)
	if err == nil && paymentRes != nil {
		return nil, commons.ErrPaymentMethodAlreadyExists
	}

	// insert payment to db
	paymentRes, err = cu.PaymentsRepository.Insert(ctx, &payment)
	if err != nil {
		return nil, err
	}

	return paymentRes, nil
}

func (cu *PaymentsUsecase) GetAll(ctx context.Context) ([]*Payment, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	// get all payments from db
	payments, err := cu.PaymentsRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (cu *PaymentsUsecase) GetById(ctx context.Context, id int) (*Payment, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	// get payment by id from db
	payment, err := cu.PaymentsRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
