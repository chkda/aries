package app

import (
	"context"

	"github.com/chkda/aries/internal/engine/database"
	"github.com/chkda/aries/internal/engine/mq"
	"github.com/chkda/aries/internal/protocols"
	"github.com/chkda/aries/pkg/mail"
)

type App struct {
	subscriber   *mq.MQSubscriber
	queryHandler *database.QueryHandler
	mailer       mail.MailSender
}

func New(
	subscriber *mq.MQSubscriber,
	queryHandler *database.QueryHandler,
	mailer mail.MailSender,
) *App {
	return &App{
		subscriber:   subscriber,
		queryHandler: queryHandler,
		mailer:       mailer,
	}
}

func (c *App) StartConsuming(ctx context.Context, queue string) {
	msgChan, err := c.subscriber.Subscribe(ctx, queue)
	if err != nil {
		return
	}

	for msg := range msgChan {
		err = c.sendMail(ctx, msg)
		status := "1"
		if err != nil {
			// log error
			status = "0"
		}
		row := &database.QueryRow{
			EventId: msg.EventId,
			Status:  status,
		}
		err = c.queryHandler.Insert(ctx, row)
		if err != nil {
			// log error
			continue
		}
	}
}

func (c *App) sendMail(ctx context.Context, msg *protocols.Message) error {
	mail := &mail.Message{
		Subject:    msg.NotificationSubject,
		Body:       msg.NotificationBody,
		Recepients: []string{msg.UserDevice},
	}

	err := c.mailer.Send(ctx, mail)
	if err != nil {
		return err
	}
	return nil
}
