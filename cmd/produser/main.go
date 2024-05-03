package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/SashaMelva/web-service-gin/internal/config"
	"github.com/SashaMelva/web-service-gin/internal/logger"
	"github.com/SashaMelva/web-service-gin/internal/server/rabbitmq"
)

func main() {
	configFile := "../../configs/"
	config := config.New(configFile)
	log := logger.New(config.Logger, "../../logs/")

	s := rabbitmq.NewConnection(log, config.ConfigRMQ)
	defer s.Stop()

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	ch, err := s.Connection.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	s.DeclareQueues(ch, "Test")
	s.PublishMessage(ctx, ch, "Test", []byte("Hi"))
}
