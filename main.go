package main

import (
	"go-ecommerce-app2/config"
	"go-ecommerce-app2/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalln("Env setup error occurred")
		return
	}

	api.StartServer(cfg)
}
