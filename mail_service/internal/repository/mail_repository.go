package repository

import (
	"crypto/tls"
	"os"

	"github.com/gookit/slog"
	gomail "gopkg.in/mail.v2"
)

func SendMail(to []string, msg []byte, subject string) error {

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", os.Getenv("MAIL"))

	// Set E-Mail receivers
	m.SetHeader("To", to...)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", string(msg))

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("MAIL"), os.Getenv("PASS"))

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		slog.Fatalf("failed to send message: \t%v", err)
		return (err)
	}

	slog.Infof("Successfully sent message")

	return nil
}
