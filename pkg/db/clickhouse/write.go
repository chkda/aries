package clickhouse

import "context"

func (c *Client) Write(ctx context.Context, query string, args ...any) error {
	_, err := c.Conn.Query(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
