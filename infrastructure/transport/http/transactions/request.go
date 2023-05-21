package transactions

import "github.com/sultanfariz/synapsis-assignment/domain/transactions"

type Transaction struct {
	UserId    int `json:"userId" form:"userId"`
	PaymentId int `json:"paymentId" form:"paymentId"`
	Products  []CheckoutList
}

type UpdateStatus struct {
	StatusId int `json:"statusId" form:"statusId" validate:"required"`
}

type CheckoutList struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

func (t Transaction) ToDomainCheckoutList() []transactions.CheckoutList {
	var request []transactions.CheckoutList
	for _, item := range t.Products {
		request = append(request, transactions.CheckoutList{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	return request
}

func (t Transaction) ToDomain() transactions.Transaction {
	return transactions.Transaction{
		UserId:    t.UserId,
		PaymentId: t.PaymentId,
		Products:  t.ToDomainCheckoutList(),
	}
}
