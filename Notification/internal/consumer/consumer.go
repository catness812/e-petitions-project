package consumer

import (
	"log"

	"github.com/catness812/e-petitions-project/Notification/internal/sender"
	"github.com/streadway/amqp"
)

type Consumer struct {
	channel *amqp.Channel
	name    string
}

func NewConsumer(channel *amqp.Channel, name string) *Consumer {
	return &Consumer{
		channel: channel,
		name:    name,
	}
}

func (c *Consumer) ConsumeAndSend() {
	msgs, err := c.channel.Consume(
		c.name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    //args
	)
	if err != nil {
		panic(err)
	}

	log.Printf("Consumer %s started", c.name)
	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			sender.SendMail(string(msg.Body))
		}
	}()

	<-forever
}
