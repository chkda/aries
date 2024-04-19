package queue

import "context"

type Message struct {
	Body  []byte
	Queue string
}

type Streamer[T any] interface {
	Publish(context.Context, *Message) error
	Subscribe(context.Context, string) (<-chan T, error)
}
