package main

import (
	"github.com/Thiago-Mota-Santos/Gotask/config"
	"github.com/Thiago-Mota-Santos/Gotask/router"
)

func main() {

	logger := config.NewLogger("")

	err := config.Init()
	if err != nil {
		logger.Errorf("error: %v", err)
	}

	router.Init()
}
