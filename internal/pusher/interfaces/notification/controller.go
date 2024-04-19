package notification

import (
	"github.com/chkda/aries/pkg/db"
	"github.com/chkda/aries/pkg/queue"
)

type Controller struct {
	Publisher queue.Publisher
	DBWriter  db.Writer
}

func New(
	publisher queue.Publisher,
	dbWriter db.Writer,
) *Controller {
	return &Controller{
		Publisher: publisher,
		DBWriter:  dbWriter,
	}
}
