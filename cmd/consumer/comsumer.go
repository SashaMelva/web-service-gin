package main

import (
	"fmt"

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

	ch, err := s.Connection.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"Test",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	log.Debug(msgs)

	if err != nil {
		log.Fatal(err)
	}

	// forver := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	// <-forever
}
