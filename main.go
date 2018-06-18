package main

import (
	"github.com/drklee3/polls-api/app"
	"github.com/drklee3/polls-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(config.GetAddr())
}
