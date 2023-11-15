package service

import (
	"regexp"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/catness812/e-petitions-project/mail_service/internal/repository"
	"github.com/gookit/slog"
)

func SendVerificationMail(msg string) error {
	var to []string

	to = append(to, strings.Split(string(msg), " ")[0])
	link := strings.Split(string(msg), " ")[1]

	err := repository.SendMail(to, formatMailMessage(link, "user-register.html"), "E-petitions verification link")
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

	to, message = getMailsAndMessage(msg)

	err := repository.SendMail(to, formatMailMessage(message, "notification.html"), "E-petitions notification")
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

func getMailsAndMessage(msg string) ([]string, string) {
	mail_pattern := "[a-zA-Z0-9.]+@[a-zA-Z0-9.-]+"

	regex := regexp.MustCompile(mail_pattern)
	mails := regex.FindAllString(msg, -1)

	message_pattern := " [a-zA-Z0-9.,!?-]+[^@]* "
	regex = regexp.MustCompile(message_pattern)
	result := regex.FindAllString(msg+" ", -1)
	return mails, result[0][1:]
}
