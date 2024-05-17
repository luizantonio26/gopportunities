package main

import (
	"github.com/luizantonio26/gopportunities/config"
	"github.com/luizantonio26/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//initialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	//initialize router
	router.Initialize()
}
