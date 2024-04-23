package clickhouse

import (
	"context"
)

func (c *Client) Write(ctx context.Context, query string, args ...any) error {
	err := c.Conn.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
