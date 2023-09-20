package main

import (
	"github.com/catness812/e-petitions-project/internal/models"
	"github.com/catness812/e-petitions-project/pkg/database/postgres"
)

func loadDatabase() {
	postgres.Connect()
	postgres.Database.AutoMigrate(&models.UserModel{})
}

func main() {
	loadDatabase()
}