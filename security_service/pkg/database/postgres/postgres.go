package postgres

import (
	"fmt"
	"github.com/catness812/e-petitions-project/security_service/internal/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	dbConfig := config.LoadConfig()
	var err error
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
		dbConfig.Database.Host,
		dbConfig.Database.DBName,
		dbConfig.Database.User,
		dbConfig.Database.Password,
		dbConfig.Database.Port,
	)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Successfully connected to the Postgres database")
	}
}
