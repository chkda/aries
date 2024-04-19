package gmail

import "net/smtp"

type Client struct {
	Auth     smtp.Auth
	Server   string
	Port     string
	Username string
}

func New(cfg *Config) *Client {
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.SMTPServer)
	return &Client{
		Auth:     auth,
		Server:   cfg.SMTPServer,
		Port:     cfg.SMTPPort,
		Username: cfg.Username,
	}
}
