package repository

import (
	"fmt"

	"github.com/catness812/e-petitions-project/petition_service/internal/config"
	"github.com/streadway/amqp"
)

type NotificationRepository struct{}

func InitNotificationRepository() *NotificationRepository {
	return &NotificationRepository{}
}

func (repo *NotificationRepository) PublishMessage(queueName string, message string) error {
	rabbitMQURL := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		config.Cfg.Broker.User,
		config.Cfg.Broker.Password,
		config.Cfg.Broker.Host,
		config.Cfg.Broker.Port,
	)
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
