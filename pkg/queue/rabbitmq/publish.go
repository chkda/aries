package rabbitmq

import (
	"context"

	"github.com/chkda/aries/pkg/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *Client) Publish(
	ctx context.Context,
	msg *queue.Message,
) error {
	channel, err := c.Conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	err = channel.PublishWithContext(
		ctx,
		"",
		msg.Queue,
		false,
		false,
		amqp.Publishing{
			Body:        msg.Body,
			ContentType: "application/json",
		},
	)
	return err
}
