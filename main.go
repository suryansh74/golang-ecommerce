package main

import (
	"log"

	"go-ecommerce-app2/config"
	"go-ecommerce-app2/internal/api"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("Env setup error occurred %v", err)
		return
	}

	api.StartServer(cfg)
}
