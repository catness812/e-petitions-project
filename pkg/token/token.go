package token

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
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


func GenerateTokenPair(user_id uint) (map[string]string, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_MINUTE_LIFESPAN"))

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(token_lifespan)).Unix()

	t, err := token.SignedString([]byte(os.Getenv("T_KEY")))

	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)

	rtoken_lifespan, err := strconv.Atoi(os.Getenv("RTOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return nil, err
	}

	rtClaims := refreshToken.Claims.(jwt.MapClaims)

	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * time.Duration(rtoken_lifespan)).Unix()

	rt, err := refreshToken.SignedString([]byte(os.Getenv("RT_KEY")))

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
