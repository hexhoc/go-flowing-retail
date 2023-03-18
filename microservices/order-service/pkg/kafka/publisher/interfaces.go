package publisher

import (
	"context"
)

type EventPublisher interface {
	Publish(ctx context.Context, msg []byte) error
	Close(ctx context.Context) error
}
