package repository

import (
	"net/smtp"
	"os"

	"github.com/gookit/slog"
)

func SendMail(to []string, msg []byte) {
	// auth := sm.SmtpAuth(os.Getenv("MAIL"), os.Getenv("PASS"))
	addr := "localhost:1025"

	err := smtp.SendMail(addr, nil, os.Getenv("MAIL"), to, msg)

	if err != nil {
		slog.Fatalf("failed to send message: \t%v", err)
		return
	}
	slog.Printf("successfully sent message")

}
