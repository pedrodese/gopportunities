package main

import (
	"github.com/pedrodese/gopportunities/config"
	"github.com/pedrodese/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger := *config.GetLogger("main")

	//Initialize configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initiliazation error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()

}
