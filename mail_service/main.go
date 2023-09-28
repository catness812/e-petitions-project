package main

import (
	"log"
	"os"

	"github.com/catness812/e-petitions-project/mail_service/internal/config"
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
	cons := controllers.NewConsumer(ch)
	cons.ConfirmationMail(q.Name)
}

func init() {
	err := godotenv.Load("./mail_service/.env")
	if err != nil {
		err = godotenv.Load("../mail_service/.env")
		if err != nil {
			log.Fatal("failed to load .env")
		}
	}

	config.LoadConfig()

	ch = rabbitMQ.ConnectAMQPDataBase(os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), config.Cfg.Rabbit.Host, config.Cfg.Rabbit.Port)
}
