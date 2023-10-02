package main

import (
	"github.com/catness812/e-petitions-project/mail_service/internal/controllers"
	"github.com/catness812/e-petitions-project/mail_service/pkg/rabbitMQ"
	"github.com/gookit/slog"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var ch *amqp.Channel

func main() {
	cons := controllers.NewConsumer(ch)
	go cons.ConfirmationMail("verify")
	cons.NotificationMail("notification")
}

func init() {
	err := godotenv.Load("./mail_service/.env")
	if err != nil {
		err = godotenv.Load("../mail_service/.env")
		if err != nil {
			slog.Fatalf("failed to load .env")
		}
	}

	ch = rabbitMQ.ConnectAMQPDataBase()
}
