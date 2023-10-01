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
	var (
		to      []string
		message string
	)

	header := "Subject: Petitionon message\nMIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n"

	for i, buf := range strings.Split(string(msg), " ") {
		if i == 0 {
			to = append(to, strings.Split(string(msg), " ")[0])
			continue
		}
		message = message + buf + " "
	}
	message = header + message

	repository.SendMail(to, []byte(message))
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
