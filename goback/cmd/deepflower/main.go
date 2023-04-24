package main

import (
	"deepflower/app"
	"deepflower/config"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	app := app.NewApp()
	if err := app.Run(cfg); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
