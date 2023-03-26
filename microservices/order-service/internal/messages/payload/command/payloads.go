package command

type RetrievePaymentCommandPayload struct {
	RefId      string  `json:"RefId"`
	CustomerId uint32  `json:"CustomerId"`
	Reason     string  `json:"Reason"`
	Amount     float64 `json:"Amount"`
}

type FetchGoodsCommandPayload struct {
	RefId  string             `json:"RefId"`
	Reason string             `json:"Reason"`
	Items  []InventoryItemDTO `json:"Items"`
}

type InventoryItemDTO struct {
	ProductId uint32 `json:"ProductId"`
	Quantity  uint32 `json:"Quantity"`
}

type ShipGoodsCommandPayload struct {
	RefId             string `json:"RefId"`
	PickId            string `json:"PickId"`
	LogisticsProvider string `json:"LogisticsProvider"`
	RecipientName     string `json:"RecipientName"`
	RecipientAddress  string `json:"RecipientAddress"`
}
