package service

import (
	"github.com/google/uuid"
	"github.com/hexhoc/payment-service/internal/messages/payload/command"
)

type PaymentInterface interface {
	Receive(command command.RetrievePaymentCommandPayload) string
}

type PaymentService struct{}

func (t *PaymentService) Receive(command command.RetrievePaymentCommandPayload) string {
	paymentId, _ := uuid.NewUUID()
	return paymentId.String()
}
