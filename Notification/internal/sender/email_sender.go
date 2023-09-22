package sender

import (
	"strings"

	"github.com/catness812/e-petitions-project/Notification/internal/service"
)

func SendMail(msg string) {
	var to []string

	to = append(to, strings.Split(string(msg), " ")[0])
	code := strings.Split(string(msg), " ")[1]

	service.SendMail(to, code)
}
