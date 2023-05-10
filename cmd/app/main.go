package main

import (
	"log"

	"github.com/staszigzag/go-custom-template/internal/app"
	"github.com/staszigzag/go-custom-template/internal/config"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
