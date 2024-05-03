package rabbitmq

import (
	"context"

	"github.com/SashaMelva/web-service-gin/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type RabbitMQServe struct {
	Connection *amqp.Connection
	log        *zap.SugaredLogger
}

func NewConnection(log *zap.SugaredLogger, config *config.ConfigRMQ) *RabbitMQServe {
	conn, err := amqp.Dial("amqp://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/")
	if err != nil {
		log.Fatal("unable to open connect to RabbitMQ server. Error:", err)
	}

	return &RabbitMQServe{
		Connection: conn,
		log:        log,
	}
}

func (s *RabbitMQServe) Stop() {
	if err := s.Connection.Close(); err != nil {
		s.log.Fatal("Server forced to shutdown: ", err)
	}
}

func (s *RabbitMQServe) DeclareQueues(ch *amqp.Channel, name string) {
	q, err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)

	s.log.Debug(q)

	if err != nil {
		s.log.Fatal(err)
	}
}

func (s *RabbitMQServe) PublishMessage(ctx context.Context, ch *amqp.Channel, nameQueues string, message []byte) {
	err := ch.PublishWithContext(
		ctx,
		"",
		nameQueues,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)

	if err != nil {
		s.log.Error(err)
	}

	s.log.Debug("Successfully Published Message to Queue")
}
