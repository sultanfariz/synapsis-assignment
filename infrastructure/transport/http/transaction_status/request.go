package transaction_status

import "github.com/sultanfariz/synapsis-assignment/domain/transaction_status"

type TransactionStatus struct {
	Status string `json:"status" form:"status" validate:"required"`
}

func (t TransactionStatus) ToDomain() transaction_status.TransactionStatus {
	return transaction_status.TransactionStatus{
		Status: t.Status,
	}
}
