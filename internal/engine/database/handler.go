package database

import (
	"context"
	"strings"

	"github.com/chkda/aries/pkg/db"
)

var columns = []string{
	"EventId",
	"Status",
}

type QueryHandler struct {
	writer db.Writer
}

type QueryRow struct {
	EventId string
	Status  string
}

func New(writer db.Writer) *QueryHandler {
	return &QueryHandler{
		writer: writer,
	}
}

func (c *QueryHandler) Insert(ctx context.Context, row *QueryRow) error {
	query := c.queryBuilder(row)
	err := c.writer.Write(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (c *QueryHandler) queryBuilder(row *QueryRow) string {
	query := "INSERT INTO " + NOTIFICATION_EVENTS_TABLE + " (" + strings.Join(columns, ",") + ") "
	values := []string{
		"'" + row.EventId + "'",
		row.Status,
	}
	query += " VALUES (" + strings.Join(values, ",") + ")"
	return query
}
