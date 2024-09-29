package configs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupPostgres(postgresUrl string) (*gorm.DB, error) {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(postgresUrl), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = db.AutoMigrate(
		&entity.Location{},
		&entity.Prediction{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}
	return db, err
}

// var setupPostgresOnce = sync.OnceValues(setupPostgres)

// func SetupPostgres() (*gorm.DB, error) {
// 	return setupPostgresOnce()
// }
