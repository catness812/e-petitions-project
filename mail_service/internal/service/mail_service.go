package service

import (
	"strings"

	"github.com/aymerick/raymond"
	"github.com/catness812/e-petitions-project/mail_service/internal/repository"
	"github.com/gookit/slog"
)

func SendVerificationMail(msg string) error {
	var to []string

	to = append(to, strings.Split(string(msg), " ")[0])
	link := strings.Split(string(msg), " ")[1]

	err := repository.SendMail(to, formatMailMessage(link, "user-register.html"))
	if err != nil {
		return err
	}

	return nil
}

func SendNotificationMail(msg string) error {
	var (
		to      []string
		message string
	)

	for i, buf := range strings.Split(string(msg), " ") {
		if i == 0 {
			to = append(to, strings.Split(string(msg), " ")[0])
			continue
		}
		message = message + buf + " "
	}

	err := repository.SendMail(to, formatMailMessage(message, "notification.html"))
	if err != nil {
		return err
	}

	return nil
}

func formatMailMessage(data string, path string) []byte {
	reg, err := raymond.ParseFile("./mail_service/templates/" + path)
	if err != nil {
		reg, err = raymond.ParseFile("../mail_service/templates/" + path)
		if err != nil {
			slog.Fatalf("failed to parse template: %v", err)
		}
	}

	ctx := map[string]interface{}{
		"data": data,
	}
	return []byte(reg.MustExec(ctx))
}
