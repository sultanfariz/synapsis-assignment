package transaction_status

import (
	"time"
)

type TransactionStatus struct {
	Id        int    `gorm:"primaryKey"`
	Status    string `gorm:"type:varchar(256);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
