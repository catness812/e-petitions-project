package rabbitMQ

import (
	"os"

	"github.com/catness812/e-petitions-project/mail_service/internal/config"
	"github.com/gookit/slog"
	"github.com/streadway/amqp"
)

func ConnectAMQPDataBase() *amqp.Channel {
	rabbit := config.LoadConfig().Rabbit
	conn, err := amqp.Dial("amqp://" + os.Getenv("RABBITMQ_USER") + ":" + os.Getenv("RABBITMQ_PASS") + "@" + rabbit.Host + ":" + rabbit.Port + "/")
	if err != nil {
		slog.Fatalf("cant start RabbitMQ: %v", err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		slog.Fatalf("cant start channal: %v", err.Error())
	}

	return ch
}
