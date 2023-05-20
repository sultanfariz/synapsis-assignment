package products

import (
	"time"

	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/categories"
)

type Product struct {
	Id          int                 `gorm:"primaryKey"`
	Name        string              `gorm:"type:varchar(512);not null"`
	Price       int                 `gorm:"type:int;not null"`
	Stock       int                 `gorm:"type:int;not null"`
	Description string              `gorm:"type:text;not null"`
	PictureUrl  string              `gorm:"type:text;not null"`
	CategoryId  int                 `gorm:"column:category_id;not null"`
	Category    categories.Category `gorm:"foreignKey:CategoryId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
