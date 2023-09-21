package rabbitMQ

import (
	"log"

	"github.com/streadway/amqp"
)

func ConnectAMQPDataBase(user string, pass string, host string, port string) *amqp.Channel {
	conn, err := amqp.Dial("amqp://" + user + ":" + pass + "@" + host + ":" + port + "/")
	if err != nil {
		log.Fatalf("cant start RabbitMQ: %v", err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("cant start channal: %v", err.Error())
	}

	return ch
}
