package messages

import (
	"fmt"

	"github.com/hexhoc/order-service/internal/service"
	"github.com/segmentio/kafka-go"
)

type MessageListener struct {
	orderService service.OrderInterface
}

func NewMessageListener(orderService service.OrderInterface) *MessageListener {
	return &MessageListener{orderService: orderService}
}

func (t *MessageListener) EventPayment(message *kafka.Message) error {
	//TODO: update order status
	fmt.Println(message)
	return nil
}

func (t *MessageListener) EventFetchGoods(message *kafka.Message) error {
	//TODO: update order status
	fmt.Println(message)
	return nil
}

func (t *MessageListener) EventShipGoods(message *kafka.Message) error {
	//TODO: update order status
	fmt.Println(message)
	return nil
}
