package event

type PaymentReceivedEventPayload struct {
	RefId     string `json:"RefId"`
	PaymentId string `json:"PaymentId"`
}
