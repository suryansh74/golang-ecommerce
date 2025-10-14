package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	ServerHost string
	ServerPort string
	Dsn        string
	AppSecret  string
}

func SetupEnv() (cfg AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}
	httpHost := os.Getenv("SERVER_HOST")
	httpPort := os.Getenv("SERVER_PORT")
	if len(httpHost) < 1 || len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variable are not set for host and port")
	}
	dsn := os.Getenv("DSN")
	appSecret := os.Getenv("APP_SECRET")

	return AppConfig{
		ServerHost: httpHost,
		ServerPort: httpPort,
		Dsn:        dsn,
		AppSecret:  appSecret,
	}, nil
}
