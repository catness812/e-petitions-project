package repository

import (
	"net/smtp"
	"os"

	"github.com/aymerick/raymond"
	"github.com/gookit/slog"
)

var reg *raymond.Template

func SendMail(to []string, link string) {
	msg := formatMessage(link)
	// auth := sm.SmtpAuth(os.Getenv("MAIL"), os.Getenv("PASS"))
	addr := "localhost:1025"

	err := smtp.SendMail(addr, nil, os.Getenv("MAIL"), to, msg)
	// err := smtp.SendMail(addr, auth, os.Getenv("MAIL"), to, msg)

	if err != nil {
		slog.Fatalf("failed to send message: \t%v", err)
		return
	}
	slog.Printf("successfully sent message")

}

func formatMessage(link string) []byte {
	ctx := map[string]interface{}{
		"link": link,
	}
	return []byte(reg.MustExec(ctx))
}

func init() {
	var err error
	reg, err = raymond.ParseFile("./mail_service/templates/user-register-link.html")
	if err != nil {
		reg, err = raymond.ParseFile("../mail_service/templates/user-register-link.html")
		if err != nil {
			slog.Fatalf("failed to parse template: %v", err)
		}
	}
}
