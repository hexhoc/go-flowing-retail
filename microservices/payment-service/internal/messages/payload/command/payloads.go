package command

type RetrievePaymentCommandPayload struct {
	RefId      string  `json:"RefId"`
	CustomerId uint32  `json:"CustomerId"`
	Reason     string  `json:"Reason"`
	Amount     float64 `json:"Amount"`
}
