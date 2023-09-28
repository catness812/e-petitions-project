package service

import (
	"strings"

	"github.com/aymerick/raymond"
	"github.com/catness812/e-petitions-project/mail_service/internal/repository"
	"github.com/gookit/slog"
)

func SendVerificationMail(msg string) {
	var to []string

	to = append(to, strings.Split(string(msg), " ")[0])
	link := strings.Split(string(msg), " ")[1]

	repository.SendMail(to, formatMessage(link))
}

func SendNotificationMail(msg string) {
	var to []string

	to = append(to, strings.Split(string(msg), " ")[0])
	message := []byte(strings.Split(string(msg), " ")[1])

	repository.SendMail(to, message)
}

func formatMessage(link string) []byte {
	reg, err := raymond.ParseFile("./mail_service/internal/templates/user-register-link.html")
	if err != nil {
		reg, err = raymond.ParseFile("../mail_service/internal/templates/user-register-link.html")
		if err != nil {
			slog.Fatalf("failed to parse template: %v", err)
		}
	}

	ctx := map[string]interface{}{
		"link": link,
	}
	return []byte(reg.MustExec(ctx))
}
