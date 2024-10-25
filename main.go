package main

import (
	"github.com/MatheusJPY/gopportunities/config"
	"github.com/MatheusJPY/gopportunities/router"
)

func main() {
	// Initialize Configs
	err := config.Init()
	if err != nil {
		panic(err)
	}
	// Initialize Router
	router.Initialize()
}
