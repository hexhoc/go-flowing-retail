package entity

import (
	"time"
)

type Order struct {
	Id         string       `json:"id"`
	CustomerId uint32       `json:"customer_id"`
	Address    string       `json:"address"`
	Status     uint32       `json:"status"`
	IsDeleted  bool         `json:"is_deleted"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	OrderItems []*OrderItem `json:"order_items"`
}
