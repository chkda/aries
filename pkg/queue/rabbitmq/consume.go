package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *Client) Subscribe(ctx context.Context, queue string) (<-chan amqp.Delivery, error) {
	channel, err := c.Conn.Channel()
	if err != nil {
		return nil, err
	}
	// defer channel.Close()

	messageChan, err := channel.ConsumeWithContext(
		ctx,
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return messageChan, nil
}
