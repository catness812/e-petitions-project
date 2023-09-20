package main

import (
	"github.com/catness812/e-petitions-project/internal/models"
	"github.com/catness812/e-petitions-project/pkg/database/postgres"
	"github.com/catness812/e-petitions-project/pkg/database/redis_repository"
)

func loadDatabase() {
	postgres.Connect()
	postgres.Database.AutoMigrate(&models.UserModel{})
	redis_repository.NewRedisDBConnection()
}

func main() {
	loadDatabase()
}
