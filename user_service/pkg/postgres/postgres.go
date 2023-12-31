package postgres

import (
	"fmt"
	"log"

	"github.com/catness812/e-petitions-project/user_service/internal/config"
	"github.com/catness812/e-petitions-project/user_service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("connected Postgres")
	}
	db.AutoMigrate(&models.User{})
	return db
}
