package main

import (
	"log"

	"github.com/ironsail/whydah-go-clean-template/config"
	"github.com/ironsail/whydah-go-clean-template/internal/app"
	"github.com/ironsail/whydah-go-clean-template/pkg/logger"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	logger.Init(cfg)

	// Run
	app.Run(cfg)
}
