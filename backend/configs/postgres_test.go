package configs

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	code := m.Run()

	os.Exit(code)
}

func TestSetupPostgres_Success(t *testing.T) {
	db, err := SetupPostgres()

	if db == nil || err != nil {
		t.Errorf("expected a valid database connection, got nil")
	}
}
