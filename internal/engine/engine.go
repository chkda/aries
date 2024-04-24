package engine

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/chkda/aries/internal/engine/app"
	"github.com/chkda/aries/internal/engine/database"
	"github.com/chkda/aries/internal/engine/interfaces/healthcheck"
	"github.com/chkda/aries/internal/engine/mq"
	"github.com/chkda/aries/pkg/db/clickhouse"
	"github.com/chkda/aries/pkg/mail/gmail"
	"github.com/chkda/aries/pkg/queue/rabbitmq"
	"github.com/labstack/echo/v4"
)

type Config struct {
	HTTPPort   string             `json:"http_port"`
	Queue      string             `json:"queue"`
	RabbitMQ   *rabbitmq.Config   `json:"rabbitmq"`
	Clickhouse *clickhouse.Config `json:"clickhouse"`
	Gmail      *gmail.Config      `json:"gmail"`
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
	mailClient := gmail.New(cfg.Gmail)
	subscriber := mq.NewSubscriber(rabbitMQClient)
	queryHandler := database.New(clickhouseClient)
	appHandler := app.New(subscriber, queryHandler, mailClient)
	ctx := context.Background()
	go func() {
		appHandler.StartConsuming(ctx, cfg.Queue)
	}()

	healthcheckController := healthcheck.New()
	serv := echo.New()
	serv.GET(healthcheckController.GetRoute(), healthcheckController.Handler)
	serv.Logger.Fatal(serv.Start(":" + cfg.HTTPPort))
}
