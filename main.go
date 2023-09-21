package main

import (
	"log"
	"notifications/internal/controllers"
	"notifications/pkg/rabbitMQ"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var ch *amqp.Channel

func main() {
	ctrl := controllers.NewEmailController(ch)
	ctrl.ReadFromQueue()
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
