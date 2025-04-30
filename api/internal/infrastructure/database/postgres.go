package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	repoModels "github.com/eijiosakabe/ai-todos/api/internal/infrastructure/repomodels"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresConnection(cfg DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	// Create a new logger which will log slow SQL statements
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 200,         // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
		},
	)

	// Open connection to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

// GetDatabaseConfigFromEnv retrieves database configuration from environment variables
func GetDatabaseConfigFromEnv() DatabaseConfig {
	return DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}

// MigrateDatabase runs database migrations
func MigrateDatabase(db *gorm.DB) error {
	// Migrate all repository models
	return db.AutoMigrate(
		&repoModels.User{},
	)
}
