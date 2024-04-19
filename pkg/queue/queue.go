package queue

import "context"

type Message struct {
	Body  []byte
	Queue string
}

type Streamer[T any] interface {
	Publish(ctx context.Context, msg *Message) error
	Subscribe(ctx context.Context, queue string) (<-chan T, error)
}
