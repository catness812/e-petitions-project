package main

import (
	"os"
	"time"

	"github.com/catness812/e-petitions-project/mail_service/internal/config"
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

	cfg := config.LoadConfig()

	defer func() {
		if r := recover(); r != nil {
			time.Sleep(time.Second * 30)
			slog.Infof("Recovered. Error:\t", r)
			ch = rabbitMQ.ConnectAMQPDataBase(os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), cfg.Rabbit.Host, cfg.Rabbit.Port)
		}
	}()

	ch = rabbitMQ.ConnectAMQPDataBase(os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), cfg.Rabbit.Host, cfg.Rabbit.Port)
}
