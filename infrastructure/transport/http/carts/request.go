package carts

import "github.com/sultanfariz/synapsis-assignment/domain/carts"

type Cart struct {
	UserId    int `json:"userId" form:"userId"`
	ProductId int `json:"productId" form:"productId"`
}

func (c Cart) ToDomain() carts.Cart {
	return carts.Cart{
		UserId:    c.UserId,
		ProductId: c.ProductId,
	}
}
