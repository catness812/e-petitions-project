package service

import (
	"log"
	"net/smtp"
	"os"

	"github.com/aymerick/raymond"
	"github.com/spf13/viper"
)

var tmp *raymond.Template

func SendMail(to []string, code string) {
	msg := formatMessage(code)
	// auth := sm.SmtpAuth(os.Getenv("MAIL"), os.Getenv("PASS"))
	addr := viper.GetString("smtpHost") + ":" + viper.GetString("smtpPort")

	err := smtp.SendMail(addr, nil, os.Getenv("MAIL"), to, msg)
	// err := smtp.SendMail(addr, auth, os.Getenv("MAIL"), to, msg)

	// handling the errors
	if err != nil {
		log.Fatalf("failed to send message: \t%v", err)
		return
	}
	log.Printf("successfully sent message")

}

func formatMessage(code string) []byte {
	subject := "Subject: Verification code\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	ctx := map[string]interface{}{
		"code": code,
	}

	return []byte(subject + mime + tmp.MustExec(ctx))
}

func init() {
	var err error
	tmp, err = raymond.ParseFile("./templates/user-register.html")
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}
}
