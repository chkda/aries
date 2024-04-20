package notification

import (
	"context"
	"net/http"

	"github.com/chkda/aries/internal/pusher/app"
	"github.com/labstack/echo/v4"
)

const (
	ROUTE = "/notifications"
)

type Controller struct {
	app *app.App
}

func (c *Controller) GetRoute() string {
	return ROUTE
}

func New(
	app *app.App,
) *Controller {
	return &Controller{
		app: app,
	}
}

func (c *Controller) Handler(e echo.Context) error {
	req := &Request{}
	resp := &Response{}
	err := e.Bind(req)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return e.JSON(http.StatusBadRequest, resp)
	}
	msg := &app.Message{
		EventId:             req.EventId,
		NotificationId:      req.NotificationId,
		NotificationSubject: req.NotificationSubject,
		NotificationBody:    req.NotificationBody,
		NotificationType:    req.NotificationType,
		UserId:              req.UserId,
		UserDevice:          req.UserDevice,
		BU:                  req.BU,
	}
	ctx := context.Background()
	err = c.app.Process(ctx, msg)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return e.JSON(http.StatusBadRequest, resp)
	}
	resp.Success = true
	resp.Message = "notification sent"
	return e.JSON(http.StatusOK, resp)
}
