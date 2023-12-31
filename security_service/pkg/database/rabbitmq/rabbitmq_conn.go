package rabbitmq

import (
	"github.com/gookit/slog"
	"github.com/streadway/amqp"
)

func ConnectAMQPDataBase(user string, pass string, host string, port string) *amqp.Channel {
	conn, err := amqp.Dial("amqp://" + user + ":" + pass + "@" + host + ":" + port + "/")
	if err != nil {
		slog.Fatalf("cant start RabbitMQ: %v", err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		slog.Fatalf("cant start channel: %v", err.Error())
	}
	return ch
}
