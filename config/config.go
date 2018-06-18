package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Username string
	Password string
	Dbname   string
	Host     string
}

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