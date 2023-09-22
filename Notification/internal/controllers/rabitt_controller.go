package controllers

import (
	"log"
	"strings"

	"github.com/catness812/e-petitions-project/Notification/internal/service/mail"
	"github.com/streadway/amqp"
)

var ch = make(chan []byte, 1)

type EmailController struct {
	ch *amqp.Channel
}

func NewEmailController(ch *amqp.Channel) *EmailController {
	return &EmailController{
		ch: ch,
	}
}

func (ctrl *EmailController) ReadFromQueue() {
	q, err := ctrl.ch.QueueDeclare("verify", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to declare registation: %v", err.Error())
	}

	msgs, err := ctrl.ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("failed to consume: %v", err)
	}

	go func() {
		for d := range msgs {
			ch <- d.Body
			go ctrl.emailSender()
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	var forever chan struct{}
	<-forever
}

func (ctrl *EmailController) emailSender() {
	var body []byte
	var to []string

	body = <-ch

	to = append(to, strings.Split(string(body), " ")[0])
	code := strings.Split(string(body), " ")[1]
	log.Printf("Received a message: \n%s\n%s\n\n", to, code)
	mail.SendMail(to, code)
}
