# Mail service

## Overview

The Mail Service is a powerful and flexible email delivery and management solution designed to simplify your email-related tasks. Whether you're sending notifications, our service is here to streamline the process and ensure your emails reach their destination reliably.

## How to Use the Project

- **Step I :** Run an RabbitMQ server.
- **Step II :** Generate an app password from mail.
- **Step III :** Configure the `*config.yml*` file
	smtp:
	  host: "smtp.example.com"
	  port: "1025"
	rabbit:
	  host: "example.com"
	  port: "5672"
- **Step IV :** Configure `*.env*` file as [example](https://github.com/catness812/e-petitions-project/blob/master/mail_service/.env.example)
	MAIL - your mail.
	PASS - generated password.
	RABBITMQ_USER - rabbitMQ user
	RABBITMQ_PASS - rabbitMQ password
- **Step V :** Run `main.go` with command  `go run main.go`

## Api

This service use rabbitMQ queues
- **verify** - this queue is used to send **OTP** to mail

- **notification** - this queue is used to send notification mails
