package main

import (
	"deepflower/app"
	"deepflower/config"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	app := app.NewApp()

	if err := app.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}

}
