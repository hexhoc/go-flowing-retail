package entity

import (
	"time"
)

type Order struct {
	Id         string       `json:"id"`
	CustomerId uint32       `json:"customer_id"`
	Address    string       `json:"address"`
	Status     string       `json:"status"`
	IsDeleted  bool         `json:"is_deleted"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	OrderItems []*OrderItem `json:"order_items"`
}

func (t *Order) GetTotalSum() float64 {
	var totalSum float64
	for _, item := range t.OrderItems {
		itemPrice, _ := item.Price.Float64()
		totalSum = totalSum + itemPrice
	}

	return totalSum
}
