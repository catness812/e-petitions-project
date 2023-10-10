package rabbitMQ

import (
	"github.com/gookit/slog"
	"github.com/streadway/amqp"
)

func ConnectAMQPDataBase(user string, pass string, host string, port string) *amqp.Channel {
	conn, err := amqp.Dial("amqp://" + user + ":" + pass + "@" + host + ":" + port + "/")
	if err != nil {
		slog.Fatalf("cant start RabbitMQ: %v", err)
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		slog.Fatalf("cant start channal: %v", err)
		panic(err)
	}

	return ch
}
