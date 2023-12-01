package main

import (
	"github.com/libdolf/gopportunities/config"
	"github.com/libdolf/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	// Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config Initialization Error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
