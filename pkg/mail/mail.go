package mail

import (
	"context"
)

type Message struct {
	Subject    string
	Body       string
	Recepients []string
}

type MailSender interface {
	Send(ctx context.Context, msg *Message) error
}
