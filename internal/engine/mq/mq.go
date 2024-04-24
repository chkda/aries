package mq

import (
	"context"
	"encoding/json"

	"github.com/chkda/aries/internal/protocols"
	"github.com/chkda/aries/pkg/queue"
)

type MQSubscriber struct {
	subscriber queue.Subscriber
}

func NewSubscriber(subsriber queue.Subscriber) *MQSubscriber {
	return &MQSubscriber{
		subscriber: subsriber,
	}
}

func (c *MQSubscriber) Subscribe(ctx context.Context, queue string) (chan *protocols.Message, error) {
	msgChan := make(chan *protocols.Message)
	subscriberChan, err := c.subscriber.Subscribe(ctx, queue)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(msgChan)
		for msgBytes := range subscriberChan {
			msg := &protocols.Message{}
			err := json.Unmarshal(msgBytes, msg)
			if err != nil {
				continue
			}
			msgChan <- msg
		}
	}()
	return msgChan, nil
}
