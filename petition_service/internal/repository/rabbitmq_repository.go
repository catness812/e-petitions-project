package repository

import (
	"fmt"
	"github.com/gookit/slog"

	"github.com/catness812/e-petitions-project/petition_service/internal/config"
	"github.com/streadway/amqp"
)

const (
	queueName = "notification"
)

type PublisherRepository struct {
	channel *amqp.Channel
}

func NewPublisherRepository() *PublisherRepository {
	slog.Info("Creating new Publisher Repository...")
	ch, err := connectToRabbit()
	if err != nil {
		slog.Errorf("Could not connect to RabbitMQ: %v", err)
	}
	return &PublisherRepository{
		channel: ch,
	}
}

func (repo *PublisherRepository) PublishMessage(email string, message string) error {
	body := fmt.Sprintf("%s %s", email, message)
	err := repo.channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// TODO move this to another layer
func connectToRabbit() (*amqp.Channel, error) {
	rabbitMQURL := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		config.Cfg.Broker.User,
		config.Cfg.Broker.Password,
		config.Cfg.Broker.Host,
		config.Cfg.Broker.Port,
	)
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}
