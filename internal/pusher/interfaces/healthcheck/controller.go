package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	ROUTE = "/healthcheck"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GetRoute() string {
	return ROUTE
}

func (c *Controller) Handler(e echo.Context) error {
	resp := &Response{
		Message: "success",
	}
	return e.JSON(http.StatusOK, resp)
}
