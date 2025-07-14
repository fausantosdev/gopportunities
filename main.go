package main

import (
	"github.com/fausantosdev/gopportunities/config"
	"github.com/fausantosdev/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("~ config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}