package categories

import (
	"time"
)

type Category struct {
	Id        int    `gorm:"primaryKey"`
	Category  string `gorm:"type:varchar(256);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
