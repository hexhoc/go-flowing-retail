package consumer

import (
	"context"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

const (
	_defaultGroupId  = "order"
	_maxRetries      = 3
	_retriesInterval = time.Second * 2
)

// MessageHeader is a key/value pair type representing headers set on records
type MessageHeader struct {
	Key   string
	Value []byte
}

// Message is a data structure representing a Kafka message
type Message struct {
	Key, Value []byte
	Topic      string
	Partition  int32
	Offset     int64
	Headers    []MessageHeader
	Timestamp  time.Time
}

// Consume reads and returns the next message from the consumer.
// The method blocks until a message becomes available, or an error occurs.
// The program may also specify a context to asynchronously cancel the blocking operation.
func Consume(ctx context.Context, addr []string, topic string, handler func(message *Message) error) {

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: addr,
		Topic:   topic,
		// initialize a new reader with the brokers and topic
		// the groupID identifies the consumer and prevents
		// it from receiving duplicate messages
		GroupID: _defaultGroupId,
	})
	defer kafkaReader.Close()

	log.Info(fmt.Sprintf("Start consuming topic: %s", kafkaReader.Config().Topic))

	go func() {
		for {
			var err error
			var msg kafka.Message
			var retryNumber = 0

			for retryNumber <= _maxRetries {
				msg, err = kafkaReader.ReadMessage(ctx)
				if err != nil {
					log.Error(err)
					retryNumber = retryNumber + 1
					time.Sleep(_retriesInterval)
					continue
				}
				break
			}

			if err != nil {
				panic(err)
			}

			log.Info(fmt.Sprintf("Received message: %s", msg.Value))

			if err := handler(messageMapping(msg)); err != nil {
				log.Printf("error consuming message, err: %#v\n", err)
				panic(err)
			}

			if err := kafkaReader.CommitMessages(context.Background(), msg); err != nil {
				panic(err)
			}
		}
	}()
}

func messageMapping(msg kafka.Message) *Message {
	var headers []MessageHeader
	if l := len(msg.Headers); l > 0 {
		headers = make([]MessageHeader, l)
		for i := range msg.Headers {
			headers[i] = MessageHeader{
				Key:   msg.Headers[i].Key,
				Value: msg.Headers[i].Value,
			}
		}
	}

	message := &Message{
		Key:       msg.Key,
		Value:     msg.Value,
		Topic:     msg.Topic,
		Partition: int32(msg.Partition),
		Offset:    msg.Offset,
		Headers:   headers,
		Timestamp: msg.Time,
	}

	return message
}
