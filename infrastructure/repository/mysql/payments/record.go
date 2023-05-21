package payments

import (
	"time"
)

type Payment struct {
	Id        int    `gorm:"primaryKey"`
	Method    string `gorm:"type:varchar(256);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
