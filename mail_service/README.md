# Description
Mail is a microservice built for the e-petition project in order to send info messages and verification link.

## Contents
- [Coneventions](#coneventions)
- [How to Use the Project](#how-to-use-the-project)
- [Api](#api)
- [Docker build](#docker-build)

## Coneventions
Mail is developed to use the rabbitMQ and will use html templates to send mails.

## How to Use the Project
- **Step I :** Run an RabbitMQ server.
- **Step II :** Generate an app password from mail.
- **Step III :** Configure the `config.yml` file
```
	smtp:
	  host: "smtp.example.com"
	  port: "1025"
	rabbit:
	  host: "example.com"
	  port: "5672"
```
- **Step IV :** Configure `.env` file as [example](https://github.com/catness812/e-petitions-project/blob/master/mail_service/.env.example)
```
	MAIL - your mail.
	PASS - generated password.
	RABBITMQ_USER - rabbitMQ user
	RABBITMQ_PASS - rabbitMQ password
```
- **Step V :** Run `main.go` with command `go run main.go`

## Api
This service use rabbitMQ queues
- **verify** - this queue is used to send **OTP** to mail
Should recive in playload message as exemple:
```example@isa.utm.md https://example.com/otp```

- **notification** - this queue is used to send notification mails
Should recive in playload message as exemple:
```example@isa.utm.md message```

## Docker build
### Command
In order to build the image necessary for the Docker compose file, run this command:
```docker build -t e-petitions-mail:1.0 .```
