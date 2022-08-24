package main

import (
	"bootstrapservice/config"
	"bootstrapservice/internal/app"
	"log"
)

func main() {
	//read config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	//run app
	app.Run(cfg)

}
