package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config stores a DBConfig
type Config struct {
	DB *DBConfig
}

// DBConfig stores database config options
type DBConfig struct {
	Username string
	Password string
	Dbname   string
	Host     string
}

// GetConfig gets the database config from environment variables
func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DB: &DBConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Dbname:   os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
		},
	}
}
