package products

import (
	"time"

	"github.com/sultanfariz/synapsis-assignment/domain/products"
)

type ProductsResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	PictureUrl  string `json:"pictureUrl"`
	CategoryId  int    `json:"categoryId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func FromDomainList(domain []*products.Product) []ProductsResponse {
	var response []ProductsResponse
	for _, item := range domain {
		response = append(response, FromDomain(*item))
	}
	return response
}

func FromDomain(domain products.Product) ProductsResponse {
	return ProductsResponse{
		Id:          domain.Id,
		Name:        domain.Name,
		Price:       domain.Price,
		Stock:       domain.Stock,
		Description: domain.Description,
		PictureUrl:  domain.PictureUrl,
		CategoryId:  domain.CategoryId,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
