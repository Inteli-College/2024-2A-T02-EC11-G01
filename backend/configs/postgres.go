package configs

import (
	"os"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupPostgres() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
