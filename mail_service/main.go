package main

import (
	"os"

	"github.com/catness812/e-petitions-project/mail_service/internal/config"
	"github.com/catness812/e-petitions-project/mail_service/internal/controllers"
	"github.com/catness812/e-petitions-project/mail_service/pkg/rabbitMQ"
	"github.com/gookit/slog"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var ch *amqp.Channel

func main() {
	q, err := ch.QueueDeclare("verify", false, false, false, false, nil)
	if err != nil {
		slog.Fatalf("failed to declare queue: %v", err)
	}
	cons := controllers.NewConsumer(ch)
	cons.ConfirmationMail(q.Name)
}

func init() {
	err := godotenv.Load("./mail_service/.env")
	if err != nil {
		err = godotenv.Load("../mail_service/.env")
		if err != nil {
			slog.Fatalf("failed to load .env")
		}
	}

	cfg := config.LoadConfig()

	ch = rabbitMQ.ConnectAMQPDataBase(os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), cfg.Rabbit.Host, cfg.Rabbit.Port)
}
