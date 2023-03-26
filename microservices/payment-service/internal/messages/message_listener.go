package messages

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hexhoc/payment-service/internal/messages/payload/command"
	"github.com/hexhoc/payment-service/internal/messages/payload/event"
	"github.com/hexhoc/payment-service/internal/service"
	"github.com/hexhoc/payment-service/pkg/kafka/consumer"
	"github.com/hexhoc/payment-service/pkg/kafka/publisher"
	log "github.com/sirupsen/logrus"
)

type MessageListener struct {
	paymentService      service.PaymentInterface
	orderEventPublisher publisher.EventPublisher
}

func NewMessageListener(paymentService service.PaymentInterface) *MessageListener {
	return &MessageListener{paymentService: paymentService}
}

func (t *MessageListener) EventPayment(message *consumer.Message) error {
	var retrievePaymentCommand command.RetrievePaymentCommandPayload
	err := json.Unmarshal(message.Value, &retrievePaymentCommand)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info(fmt.Sprintf("Retrieve payment: %d for %s", retrievePaymentCommand.Amount, retrievePaymentCommand.RefId))

	// Processing. Long operation
	paymentId := t.paymentService.Receive(retrievePaymentCommand)
	paymentReceived := &event.PaymentReceivedEventPayload{
		RefId:     retrievePaymentCommand.RefId,
		PaymentId: paymentId,
	}

	messageJson, _ := json.Marshal(paymentReceived)
	t.orderEventPublisher.Publish(context.Background(), messageJson)

	return nil
}
