package postgres

import (
	"fmt"
	config "github.com/catness812/e-petitions-project/file_service/internal"
	"github.com/catness812/e-petitions-project/file_service/internal/model"
	"github.com/gookit/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dbConfig *config.Config) *gorm.DB {
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
		slog.Info("Successfully connected to the Postgres database")
	}

	autoMigrate(database)

	return database
}

func autoMigrate(database *gorm.DB) {

	if err := database.AutoMigrate(&model.File{}); err != nil {
		slog.Fatalf("failed to automigrate file model: %v", err)
	}
	if err := database.AutoMigrate(&model.User{}); err != nil {
		slog.Fatalf("failed to automigrate chunk model: %v", err)
	}
	if err := database.AutoMigrate(&model.UserPhoto{}); err != nil {
		slog.Fatalf("failed to automigrate chunk model: %vd", err)
	}
}
