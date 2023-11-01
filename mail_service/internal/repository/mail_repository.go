package repository

import (
	"net/smtp"
	"os"

	"github.com/catness812/e-petitions-project/mail_service/internal/config"
	sm "github.com/catness812/e-petitions-project/mail_service/pkg/smtp"
	"github.com/gookit/slog"
)

func SendMail(to []string, msg []byte) error {
	s := config.LoadConfig().Smtp
	auth := sm.SmtpAuth(os.Getenv("MAIL"), os.Getenv("PASS"))
	addr := s.Host + ":" + s.Port

	err := smtp.SendMail(addr, auth, os.Getenv("MAIL"), to, msg)

	if err != nil {
		slog.Fatalf("failed to send message: \t%v", err)
		return err
	}
	slog.Infof("successfully sent message")
	return nil
}
