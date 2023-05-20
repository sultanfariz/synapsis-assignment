package categories

import "github.com/sultanfariz/synapsis-assignment/domain/categories"

type Category struct {
	Category string `json:"category" form:"category" validate:"required"`
}

func (c Category) ToDomain() categories.Category {
	return categories.Category{
		Category: c.Category,
	}
}
