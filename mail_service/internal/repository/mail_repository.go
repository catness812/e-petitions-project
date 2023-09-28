package repository

import (
	"log"
	"net/smtp"
	"os"

	"github.com/aymerick/raymond"
)

var reg *raymond.Template

func SendMail(to []string, link string) {
	msg := formatMessage(link)
	// auth := sm.SmtpAuth(os.Getenv("MAIL"), os.Getenv("PASS"))
	addr := "localhost:1025"

	err := smtp.SendMail(addr, nil, os.Getenv("MAIL"), to, msg)
	// err := smtp.SendMail(addr, auth, os.Getenv("MAIL"), to, msg)

	if err != nil {
		log.Fatalf("failed to send message: \t%v", err)
		return
	}
	log.Printf("successfully sent message")

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
			log.Fatalf("failed to parse template: %v", err)
		}
	}
}
