package main

import (
	"log"
	"os"

	"github.com/catness812/e-petitions-project/email_service/internal/controllers"
	"github.com/catness812/e-petitions-project/email_service/pkg/rabbitMQ"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
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

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")

	err = viper.ReadInConfig()
	if err != nil {

		log.Fatalf("failed to read config")
	}

	ch = rabbitMQ.ConnectAMQPDataBase(os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), viper.GetString("rabbitHost"), viper.GetString("rabbitPort"))
}
