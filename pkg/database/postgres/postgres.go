package postgres

import (
	"fmt"
	"log"

	"github.com/catness812/e-petitions-project/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	dbConfig := config.Cfg.Database
	var err error
	log.Printf("%s", dbConfig);
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable",
    dbConfig.Host,
    dbConfig.DBName,
    dbConfig.User,
    dbConfig.Password,
    dbConfig.Port,
)

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Successfully connected to the Postgres database")
	}
}