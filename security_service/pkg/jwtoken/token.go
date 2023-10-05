package jwtoken

import (
	"errors"
	"github.com/catness812/e-petitions-project/security_service/internal/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gookit/slog"
)

func verifyToken(t string) (*jwt.Token, error) {
	keyConfig := config.LoadConfig()
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(keyConfig.Token.TKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func IsTokenValid(t string) (string, error) {
	token, err := verifyToken(t)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", err
	}

	emailField, ok := claims["userEmail"]
	if !ok {
		slog.Info("Claim was not found in payload")
		return "", err
	}
	email, ok := emailField.(string)
	if !ok {
		slog.Info("claim does not contain email")
	}
	return email, err
}
