package rabbitmq

import (
	"context"
)

func (c *Client) Subscribe(ctx context.Context, queue string) (chan []byte, error) {
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

	bodyChan := make(chan []byte)
	go func() {
		defer close(bodyChan)
		defer channel.Close()
		for msg := range messageChan {
			bodyChan <- msg.Body
		}
	}()

	return bodyChan, nil
}
