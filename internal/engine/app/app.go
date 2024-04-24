package app

import "github.com/chkda/aries/internal/engine/mq"

type App struct {
	subscriber *mq.MQSubscriber
}

func New(subscriber *mq.MQSubscriber) *App {
	return &App{
		subscriber: subscriber,
	}
}
