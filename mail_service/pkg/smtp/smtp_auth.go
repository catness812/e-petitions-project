package smtp

import (
	"errors"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

func auth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unkown fromServer")
		}
	}
	return nil, nil
}

func SmtpAuth(from string, pass string) smtp.Auth {
	if from == "" || pass == "" {
		return nil
	}
	auth := auth(from, pass)
	return auth
}
