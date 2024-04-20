package mq

import (
	"context"
	"encoding/json"

	"github.com/chkda/aries/internal/protocols"
	"github.com/chkda/aries/pkg/queue"
)

type MQPublisher struct {
	publisher queue.Publisher
}

func New(publisher queue.Publisher) *MQPublisher {
	return &MQPublisher{
		publisher: publisher,
	}
}

func (c *MQPublisher) SendMessage(
	ctx context.Context,
	msg *protocols.Message,
) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	msgBody := &queue.Message{
		Body:  data,
		Queue: protocols.NOTFICATION_EVENTS_QUEUE,
	}
	err = c.publisher.Publish(ctx, msgBody)
	if err != nil {
		return err
	}
	return nil
}
