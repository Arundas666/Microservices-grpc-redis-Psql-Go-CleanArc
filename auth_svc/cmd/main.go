package main

import (
	"auth_svc/pkg/config"
	"auth_svc/pkg/di"
	"log"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config", err.Error())
	}
	server, err1 := di.InitializeServe(c)
	if err1 != nil {
		log.Fatal("Failed to init server", err1.Error())
	}

	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server")
	}
}
