package carts

import (
	"time"

	"github.com/sultanfariz/synapsis-assignment/domain/carts"
)

type CartsResponse struct {
	Id        int `json:"id"`
	UserId    int `json:"userId"`
	ProductId int `json:"productId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomainList(domain []*carts.Cart) []CartsResponse {
	var response []CartsResponse
	for _, item := range domain {
		response = append(response, FromDomain(*item))
	}
	return response
}

func FromDomain(domain carts.Cart) CartsResponse {
	return CartsResponse{
		Id:        domain.Id,
		UserId:    domain.UserId,
		ProductId: domain.ProductId,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
