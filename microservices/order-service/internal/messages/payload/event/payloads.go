package event

type GoodsFetchedEventPayload struct {
	RefId  string `json:"RefId"`
	PickId string `json:"PickId"`
}

type GoodsShippedEventPayload struct {
	RefId      string `json:"RefId"`
	ShipmentId string `json:"ShipmentId"`
}

type OrderCompletedEventPayload struct {
	OrderId string `json:"OrderId"`
}

type PaymentReceivedEventPayload struct {
	RefId     string `json:"RefId"`
	PaymentId string `json:"PaymentId"`
}
