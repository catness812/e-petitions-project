package mail

import (
	"log"
	"net/smtp"
	sm "notifications/pkg/smtp"
	"os"
)

func SendMail(to []string, code string) {
	subject := "Subject: Test mail\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(subject + mime + Message(code))

	auth := sm.SmtpAuth(os.Getenv("MAIL"), os.Getenv("PASS"))

	err := smtp.SendMail("smtp-mail.outlook.com"+":587", auth, "popov.adrian@outlook.com", to, msg)

	// handling the errors
	if err != nil {
		log.Fatalf("failed to send message: \t%v", err)
	}
	log.Printf("successfully sent message")
}
