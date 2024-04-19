package gmail

import (
	"context"
	"net/smtp"

	"github.com/chkda/aries/pkg/mail"
)

func (c *Client) Send(ctx context.Context, msg *mail.Message) error {
	err := smtp.SendMail(
		c.Server+":"+c.Port,
		c.Auth,
		c.Username,
		msg.Recepients,
		[]byte(msg.Subject+msg.Body),
	)
	if err != nil {
		return err
	}
	return nil
}
