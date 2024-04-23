package pusher

import (
	"encoding/json"
	"io"
	"os"

	"github.com/chkda/aries/internal/pusher/app"
	"github.com/chkda/aries/internal/pusher/database"
	"github.com/chkda/aries/internal/pusher/interfaces/healthcheck"
	"github.com/chkda/aries/internal/pusher/interfaces/notification"
	"github.com/chkda/aries/internal/pusher/mq"
	"github.com/chkda/aries/pkg/db/clickhouse"
	"github.com/chkda/aries/pkg/queue/rabbitmq"
	"github.com/labstack/echo/v4"
)

type Config struct {
	HTTPPort   string             `json:"http_port"`
	RabbitMQ   *rabbitmq.Config   `json:"rabbitmq"`
	Clickhouse *clickhouse.Config `json:"clickhouse"`
}

func Start(config string) {
	file, err := os.Open(config)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	cfg := &Config{}
	err = json.Unmarshal(fileBytes, cfg)
	if err != nil {
		panic(err)
	}
	rabbitMQClient, err := rabbitmq.New(cfg.RabbitMQ)
	if err != nil {
		panic(err)
	}
	clickhouseClient, err := clickhouse.New(cfg.Clickhouse)
	if err != nil {
		panic(err)
	}
	publisher := mq.NewPublisher(rabbitMQClient)
	queryHandler := database.New(clickhouseClient)
	appHandler := app.New(queryHandler, publisher)

	notificationController := notification.New(appHandler)
	healthcheckController := healthcheck.New()
	serv := echo.New()
	serv.GET(healthcheckController.GetRoute(), healthcheckController.Handler)
	serv.POST(notificationController.GetRoute(), notificationController.Handler)
	serv.Logger.Fatal(serv.Start(":" + cfg.HTTPPort))
}
