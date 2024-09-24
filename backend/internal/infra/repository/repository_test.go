package repository

import (
	"flag"
	"log"
	"os"
	"testing"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/configs"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

var db *gorm.DB
var locationRepo *LocationRepositoryGorm
var predictionRepo *PredictionRepositoryGorm

func TestMain(m *testing.M) {
	loadRepo := flag.String("loadrepo", "default", "Which repository must load")

	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	db, err = configs.SetupPostgres()
	if err != nil {
		log.Fatalf("Error setting up database connection: %v", err)
	}

	switch *loadRepo {
	case "location":
		locationRepo = NewLocationRepositoryGorm(db)
	case "prediction":
		log.Println("Loading only prediction")
		predictionRepo = NewPredictionRepositoryGorm(db)
	default:
		locationRepo = NewLocationRepositoryGorm(db)
		predictionRepo = NewPredictionRepositoryGorm(db)

	}

	code := m.Run()

	os.Exit(code)
}
