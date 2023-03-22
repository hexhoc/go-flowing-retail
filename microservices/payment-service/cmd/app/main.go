package main

import (
	"fmt"

	"github.com/hexhoc/payment-service/config"
	"github.com/hexhoc/payment-service/internal/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(fmt.Errorf("config error: %s", err))
	}

	// Run
	app.Run(cfg)
}
