package clickhouse

import (
	clk "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Client struct {
	Conn driver.Conn
}

func New(cfg *Config) (*Client, error) {
	conn, err := clk.Open(
		&clk.Options{
			Addr: []string{cfg.Host},
			Auth: clk.Auth{
				Database: cfg.Database,
				Username: cfg.Username,
				Password: cfg.Password,
			},
		},
	)

	if err != nil {
		return nil, err
	}
	return &Client{
		Conn: conn,
	}, nil
}
