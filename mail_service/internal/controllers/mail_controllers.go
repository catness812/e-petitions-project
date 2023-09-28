package controllers

import (
	"log"

	"github.com/catness812/e-petitions-project/mail_service/internal/service"
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
	msgs, err := c.channel.Consume(
		name,  // queue
		"",    // consumer
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   //args
	)
	if err != nil {
		panic(err)
	}

	log.Printf("Consumer %s started", name)
	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			service.SendMail(string(msg.Body))
		}
	}()

	<-forever
}

func (c *Consumer) NotificationMail(name string) {
	msgs, err := c.channel.Consume(
		name,  // queue
		"",    // consumer
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   //args
	)
	if err != nil {
		panic(err)
	}

	log.Printf("Consumer %s started", name)
	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			service.SendMail(string(msg.Body))
		}
	}()

	<-forever
}
