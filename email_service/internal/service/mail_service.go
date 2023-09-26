package service

import (
	"strings"

	"github.com/catness812/e-petitions-project/email_service/internal/repository"
)

func SendMail(msg string) {
	var to []string

	to = append(to, strings.Split(string(msg), " ")[0])
	code := strings.Split(string(msg), " ")[1]

	repository.SendMail(to, code)
}
