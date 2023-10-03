package jwtoken

import (
	"fmt"

	"github.com/catness812/e-petitions-project/security_service/internal/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gookit/slog"
)

func verifyToken(t string) (*jwt.Token, error) {
	keyConfig := config.LoadConfig();
	
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
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

	emailField, found := claims["userEmail"]
	if !found {
		slog.Printf("Claim was not found in payload")
		return "", err
	}

	email := fmt.Sprintf("%v", emailField)

	return email, err
}

