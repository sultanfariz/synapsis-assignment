package checkout_list

import (
	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/products"
	"github.com/sultanfariz/synapsis-assignment/infrastructure/repository/mysql/transactions"
)

type CheckoutList struct {
	Id            int                      `gorm:"primaryKey"`
	TransactionId int                      `gorm:"column:transaction_id;not null"`
	Transaction   transactions.Transaction `gorm:"foreignKey:TransactionId"`
	ProductId     int                      `gorm:"column:product_id;not null"`
	Product       products.Product         `gorm:"foreignKey:ProductId"`
	Quantity      int                      `gorm:"column:quantity;not null"`
}
