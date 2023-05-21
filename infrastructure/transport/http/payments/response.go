package payments

import (
	"time"

	"github.com/sultanfariz/synapsis-assignment/domain/payments"
)

type PaymentsResponse struct {
	Id        int    `json:"id"`
	Method    string `json:"method"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomainList(domain []*payments.Payment) []PaymentsResponse {
	var response []PaymentsResponse
	for _, item := range domain {
		response = append(response, FromDomain(*item))
	}
	return response
}

func FromDomain(domain payments.Payment) PaymentsResponse {
	return PaymentsResponse{
		Id:        domain.Id,
		Method:    domain.Method,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
