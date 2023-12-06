package controllers

import (
	"time"

	"github.com/catness812/e-petitions-project/mail_service/internal/service"
	"github.com/gookit/slog"
	"github.com/streadway/amqp"
)

type Consumer struct {
	channel *amqp.Channel
}

func NewConsumer(channel *amqp.Channel) *Consumer {
	return &Consumer{
		channel: channel,
	}
}

func (c *Consumer) ConfirmationMail(name string) {
	q, err := c.channel.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		slog.Fatalf("failed to declare queue: %v", err)
	}

	msgs, err := c.channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		slog.Panic(err)
	}

	slog.Infof("Consumer %s started", name)
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			if err := service.SendVerificationMail(string(msg.Body)); err != nil {
				time.Sleep(time.Minute)
				c.channel.Publish("",
					q.Name,
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        msg.Body,
					},
				)
			}
		}
	}()

	<-forever
}

func (c *Consumer) NotificationMail(name string) {
	q, err := c.channel.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		slog.Fatalf("failed to declare queue: %v", err)
	}

	msgs, err := c.channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		slog.Panic(err)
	}

	slog.Infof("Consumer %s started", name)
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			if err := service.SendNotificationMail(string(msg.Body)); err != nil {
				time.Sleep(time.Minute)
				c.channel.Publish("",
					q.Name,
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        msg.Body,
					},
				)
			}
		}
	}()

	<-forever
}
