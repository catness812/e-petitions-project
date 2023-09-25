package jwtoken

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
)

func verifyToken(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("RT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func IsTokenValid(t string) (jwt.MapClaims, error) {
	token, err := verifyToken(t)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, err
	}
	return claims, err
}
