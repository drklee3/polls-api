package main

import (
	"github.com/drklee3/polls-api/api/app"
	"github.com/drklee3/polls-api/api/config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(config.GetAddr())
}
