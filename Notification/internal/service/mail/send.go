package mail

import (
	"log"
	"net/smtp"
	"os"

	sm "github.com/catness812/e-petitions-project/Notification/pkg/smtp"
	"github.com/spf13/viper"
)

func SendMail(to []string, code string) {
	subject := "Subject: Verification code\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(subject + mime + Message(code))

	auth := sm.SmtpAuth(os.Getenv("MAIL"), os.Getenv("PASS"))

	addr := viper.GetString("smtpHost") + ":" + viper.GetString("smtpPort")

	err := smtp.SendMail(addr, auth, os.Getenv("MAIL"), to, msg)

	// handling the errors
	if err != nil {
		log.Fatalf("failed to send message: \t%v", err)
		return
	}
	log.Printf("successfully sent message")

}
