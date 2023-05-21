package transaction_status

import (
	"time"

	"github.com/sultanfariz/synapsis-assignment/domain/transaction_status"
)

type TransactionStatusResponse struct {
	Id        int    `json:"id"`
	Status    string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomainList(domain []*transaction_status.TransactionStatus) []TransactionStatusResponse {
	var response []TransactionStatusResponse
	for _, item := range domain {
		response = append(response, FromDomain(*item))
	}
	return response
}

func FromDomain(domain transaction_status.TransactionStatus) TransactionStatusResponse {
	return TransactionStatusResponse{
		Id:        domain.Id,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
