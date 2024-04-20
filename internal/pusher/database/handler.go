package database

import (
	"context"
	"strings"

	"github.com/chkda/aries/pkg/db"
)

var columns = []string{
	"EventId",
	"NotificationId",
	"Subject",
	"Body",
	"UserId",
	"UserDevice",
	"NotificationType",
	"BU",
}

type QueryHandler struct {
	writer db.Writer
}

type QueryRow struct {
	EventId             string
	NotificationId      string
	NotificationSubject string
	NotificationBody    string
	UserId              string
	UserDevice          string
	NotificationType    string
	BU                  string
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
		row.EventId,
		row.NotificationId,
		row.NotificationSubject,
		row.NotificationBody,
		row.UserId,
		row.UserDevice,
		row.NotificationType,
		row.BU,
	}
	query += " VALUES " + strings.Join(values, ",")
	return query
}
