package postgres

import (
	"fmt"
	"github.com/catness812/e-petitions-project/security_service/internal/config"
	"github.com/gookit/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dbConfig := config.LoadConfig()
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
		dbConfig.Database.Host,
		dbConfig.Database.DBName,
		dbConfig.Database.User,
		dbConfig.Database.Password,
		dbConfig.Database.Port,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Fatal(err)
	} else {
		slog.Println("Successfully connected to the Postgres database")
	}
	return database
}
