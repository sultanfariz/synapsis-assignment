package transactions

import (
	"time"

	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/payments"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/transaction_status"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/users"
)

type Transaction struct {
	Id        int                                  `gorm:"primaryKey"`
	UserId    int                                  `gorm:"column:user_id;not null"`
	User      users.User                           `gorm:"foreignKey:UserId"`
	StatusId  int                                  `gorm:"column:status_id;not null"`
	Status    transaction_status.TransactionStatus `gorm:"foreignKey:StatusId"`
	PaymentId int                                  `gorm:"column:payment_id;not null"`
	Payment   payments.Payment                     `gorm:"foreignKey:PaymentId"`
	TotalCost int                                  `gorm:"column:total_cost;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
