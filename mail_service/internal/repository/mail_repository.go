package repository

import (
	"net/smtp"
	"os"

	"github.com/catness812/e-petitions-project/mail_service/internal/config"
	"github.com/gookit/slog"
)

func SendMail(to []string, msg []byte) {
	s := config.LoadConfig().Smtp
	// auth := sm.SmtpAuth(os.Getenv("MAIL"), os.Getenv("PASS"))
	addr := s.Host + ":" + s.Port

	err := smtp.SendMail(addr, nil, os.Getenv("MAIL"), to, msg)

	if err != nil {
		slog.Fatalf("failed to send message: \t%v", err)
		return
	}
	slog.Printf("successfully sent message")

}
