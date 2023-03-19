package publisher

import (
	"context"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const (
	// RequireAllReplicas means that ALL nodes in the replica-set must to confirm the write for a write
	// to be considered durable.
	_requireAllReplicas = -1

	// DefaultWriteTimeout is how much to wait for a write to go through.
	_defaultWriteTimeout = 30 * time.Second

	// DefaultReadTimeout is how much to wait for reads.
	_defaultReadTimeout = 30 * time.Second

	// DefaultMaxRetryAttempts is how many times to retry a failed operation.
	_defaultMaxRetryAttempts = 3
)

type EventPublisher interface {
	Publish(ctx context.Context, msg []byte) error
	Close(ctx context.Context) error
}

type Producer struct {
	writer *kafka.Writer
}

// Addr:     kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
// Topic:   "topic-A",
// Balancer: &kafka.LeastBytes{},
func NewPublisher(addr []string, topic string) *Producer {

	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP(addr...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},

		WriteTimeout: _defaultWriteTimeout,
		ReadTimeout:  _defaultReadTimeout,

		// Use a safe default for durability.
		RequiredAcks: _requireAllReplicas,
		MaxAttempts:  _defaultMaxRetryAttempts,
	}

	return &Producer{w}
}

func (p *Producer) Publish(ctx context.Context, msg []byte) error {
	kafkaMessage := kafka.Message{
		Value: msg,
	}
	return p.writer.WriteMessages(ctx, kafkaMessage)
}

// Close tries to close the producer, but it will return sooner if the context is canceled.
// A routine in background will still try to close the producer since the underlying library does not support
// contexts on Close().
func (p *Producer) Close(ctx context.Context) error {
	done := make(chan error, 1)
	go func() {
		if p.writer != nil {
			done <- p.writer.Close()
		}
		close(done)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}
