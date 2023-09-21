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
	var err error

	dsn := fmt.Sprintf(`host=%s
	dbname=%s
	user=%s
	password=%s
	port=%d
	sslmode=disable`,
		config.Cfg.Database.Host,
		config.Cfg.Database.DBName,
		config.Cfg.Database.User,
		config.Cfg.Database.Password,
		config.Cfg.Database.Port,
	)

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Successfully connected to the Postgres database")
	}
}
