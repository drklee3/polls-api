package config

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
	return &Config{
		DB: &DBConfig{
			Username: "guest",
			Password: "Guest0000!",
			Dbname:     "todoapp",
			Host:     "localhost",
		},
	}
}