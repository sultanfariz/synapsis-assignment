package payments

import "github.com/sultanfariz/synapsis-assignment/domain/payments"

type Payment struct {
	Method string `json:"method" form:"method" validate:"required"`
}

func (p Payment) ToDomain() payments.Payment {
	return payments.Payment{
		Method: p.Method,
	}
}
