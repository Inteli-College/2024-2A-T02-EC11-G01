package configs

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupPostgres() (*gorm.DB, error) {
	postgresUrl, isSet := os.LookupEnv("POSTGRES_URL")
	if !isSet {
		log.Fatalf("POSTGRES_URL is not set")
	}

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

	return db, err
}

var setupPostgresOnce = sync.OnceValues(setupPostgres)

func SetupPostgres() (*gorm.DB, error) {
	return setupPostgresOnce()
}
