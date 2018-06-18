package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config stores a DBConfig
type Config struct {
	DB        *DBConfig
	Interface string
	Port      string
}

// GetAddr gets the full network address / port to listen on
func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%s", c.Interface, c.Port)
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

	config := &Config{
		DB: &DBConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Dbname:   os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
		},
		Interface: os.Getenv("INTERFACE"),
		Port:      os.Getenv("PORT"),
	}

	// check if interface given, use localhost as default
	if len(config.Interface) == 0 {
		config.Interface = "127.0.0.1"
	}

	// check if port given, use 3000 as default
	if len(config.Port) == 0 {
		config.Port = "3000"
	}

	return config
}
