package rabbitMQ

import (
	"github.com/streadway/amqp"
)

func ConnectAMQPDataBase(user string, pass string, host string, port string) *amqp.Channel {
	conn, err := amqp.Dial("amqp://" + user + ":" + pass + "@" + host + ":" + port + "/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch
}
