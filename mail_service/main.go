package main

import (
	"log"
	"os"

	"github.com/catness812/e-petitions-project/mail_service/internal/controllers"
	"github.com/catness812/e-petitions-project/mail_service/pkg/rabbitMQ"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var ch *amqp.Channel

func main() {
	q, err := ch.QueueDeclare("verify", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
	}
	cons := controllers.NewConsumer(ch, q.Name)
	cons.ConsumeAndSend()
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ch = rabbitMQ.ConnectAMQPDataBase(os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), "localhost", "5672")
}
