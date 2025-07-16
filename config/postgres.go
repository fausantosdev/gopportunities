package config

import (
	"os"

	"github.com/fausantosdev/gopportunities/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	err := godotenv.Load()
	if err != nil {
		logger.Errorf("x error loading .env file %e", err)
	}

	connStr := os.Getenv("DATABASE_URL")

	// Create database and connect
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Errorf("x database connection error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("x database automigration error: %v", err)
		return nil, err
	}

	return db, nil
}