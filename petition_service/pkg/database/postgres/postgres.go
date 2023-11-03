package postgres

import (
	"fmt"

	"github.com/catness812/e-petitions-project/petition_service/internal/config"
	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
	"github.com/gookit/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDatabase() *gorm.DB {
	db := connect()
	err := db.AutoMigrate(&models.Petition{})
	if err != nil {
		slog.Error(err)
	}
	err = db.AutoMigrate(&models.Status{})
	if err != nil {
		slog.Error(err)
	}
	err = db.AutoMigrate(&models.Vote{})
	if err != nil {
		slog.Error(err)
	}
	err = db.Exec("ALTER TABLE votes ALTER COLUMN petition_id TYPE VARCHAR(255)").Error
	if err != nil {
		slog.Error("Error altering column type: ", err)
	}

	seedStatuses(db)
	return db
}

func connect() *gorm.DB {
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

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		slog.Error(err)
	} else {
		slog.Info("Successfully connected to the Postgres database")
	}

	return database
}

func Paginate(pagination util.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}

// seeds Statuses in case there are none in the db
func seedStatuses(db *gorm.DB) {
	var count int64
	db.Model(&models.Status{}).Count(&count)
	if count == 0 {
		for _, status := range models.StatusSeedData {
			db.Create(&status)
		}
	}
}
