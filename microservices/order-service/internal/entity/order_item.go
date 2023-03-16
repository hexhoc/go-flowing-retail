package entity

import (
	"math/big"
)

type OrderItem struct {
	Id        string    `json:"id"`
	OrderId   string    `json:"order_id"`
	ProductId uint32    `json:"product_id"`
	Quantity  uint32    `json:"quantity"`
	Price     big.Float `json:"price"`
}
