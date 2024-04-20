package app

import (
	"context"

	"github.com/chkda/aries/internal/protocols"
	"github.com/chkda/aries/internal/pusher/database"
	"github.com/chkda/aries/internal/pusher/mq"
)

type Message struct {
	EventId             string
	NotificationId      string
	NotificationSubject string
	NotificationBody    string
	UserId              string
	UserDevice          string
	NotificationType    string
	BU                  string
}

type App struct {
	queryHandler *database.QueryHandler
	publisher    *mq.MQPublisher
}

func New(
	queryhandler *database.QueryHandler,
	publisher *mq.MQPublisher,
) *App {
	return &App{
		queryHandler: queryhandler,
		publisher:    publisher,
	}
}

func (c *App) Process(ctx context.Context, msg *Message) error {
	queryRow := &database.QueryRow{
		EventId:             msg.EventId,
		NotificationId:      msg.NotificationId,
		NotificationSubject: msg.NotificationSubject,
		NotificationBody:    msg.NotificationBody,
		NotificationType:    msg.NotificationType,
		UserId:              msg.UserId,
		UserDevice:          msg.UserDevice,
		BU:                  msg.BU,
	}
	err := c.queryHandler.Insert(ctx, queryRow)
	if err != nil {
		return err
	}
	err = c.publisher.SendMessage(ctx, &protocols.Message{
		EventId:             msg.EventId,
		NotificationId:      msg.NotificationId,
		NotificationSubject: msg.NotificationSubject,
		NotificationBody:    msg.NotificationBody,
		NotificationType:    msg.NotificationType,
		UserId:              msg.UserId,
		UserDevice:          msg.UserDevice,
		BU:                  msg.BU,
	})
	if err != nil {
		return err
	}
	return nil
}
