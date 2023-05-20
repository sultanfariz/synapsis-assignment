package products

import "github.com/sultanfariz/synapsis-assignment/domain/products"

type Product struct {
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
	Description string `json:"description" form:"description"`
	PictureUrl  string `json:"pictureUrl" form:"pictureUrl"`
	CategoryId  int    `json:"categoryId" form:"categoryId"`
}

func (p Product) ToDomain() products.Product {
	return products.Product{
		Name:        p.Name,
		Price:       p.Price,
		Stock:       p.Stock,
		Description: p.Description,
		PictureUrl:  p.PictureUrl,
		CategoryId:  p.CategoryId,
	}
}
