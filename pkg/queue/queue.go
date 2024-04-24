package queue

import "context"

type Message struct {
	Body  []byte
	Queue string
}

type Publisher interface {
	Publish(ctx context.Context, msg *Message) error
}

type Subscriber interface {
	Subscribe(ctx context.Context, queue string) (chan []byte, error)
}
