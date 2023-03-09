package main

import (
	"fmt"
	"github.com/hexhoc/product-service/config"
	"github.com/hexhoc/product-service/internal/app"
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
