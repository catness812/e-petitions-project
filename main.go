package main

import (
	"log"
	"notifications/internal/service/mail"

	"github.com/joho/godotenv"
)

func main() {
	to := []string{
		"popov.adrian@outlook.com",
	}

	mail.SendMail(to, "454665")
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
