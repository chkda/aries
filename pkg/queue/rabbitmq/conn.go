package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	Conn *amqp.Connection
}

func New(cfg *Config) (*Client, error) {
	host := fmt.Sprintf("amqp://%s:%s@%s/", cfg.Username, cfg.Password, cfg.Host)
	conn, err := amqp.Dial(host)
	if err != nil {
		return nil, err
	}

	return &Client{
		Conn: conn,
	}, nil
}
