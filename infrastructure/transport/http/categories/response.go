package categories

import (
	"time"

	"github.com/sultanfariz/synapsis-assignment/domain/categories"
)

type CategoriesResponse struct {
	Id        int    `json:"id"`
	Category  string `json:"category"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomainList(domain []*categories.Category) []CategoriesResponse {
	var response []CategoriesResponse
	for _, item := range domain {
		response = append(response, FromDomain(*item))
	}
	return response
}

func FromDomain(domain categories.Category) CategoriesResponse {
	return CategoriesResponse{
		Id:        domain.Id,
		Category:  domain.Category,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
