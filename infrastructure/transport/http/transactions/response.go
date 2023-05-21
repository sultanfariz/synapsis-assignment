package transactions

import (
	"time"

	"github.com/sultanfariz/synapsis-assignment/domain/transactions"
)

type TransactionsResponse struct {
	Id        int            `json:"id"`
	UserId    int            `json:"userId"`
	StatusId  int            `json:"statusId"`
	PaymentId int            `json:"paymentId"`
	TotalCost int            `json:"totalCost"`
	Products  []CheckoutList `json:"products"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CheckoutListResponse struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

func FromDomainList(domain []*transactions.Transaction) []TransactionsResponse {
	var response []TransactionsResponse
	for _, item := range domain {
		response = append(response, FromDomain(*item))
	}
	return response
}

func FromDomainCheckoutList(domain []transactions.CheckoutList) []CheckoutList {
	var response []CheckoutList
	for _, item := range domain {
		response = append(response, CheckoutList{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	return response
}

func FromDomain(domain transactions.Transaction) TransactionsResponse {
	return TransactionsResponse{
		Id:        domain.Id,
		UserId:    domain.UserId,
		StatusId:  domain.StatusId,
		PaymentId: domain.PaymentId,
		TotalCost: domain.TotalCost,
		Products:  FromDomainCheckoutList(domain.Products),
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
