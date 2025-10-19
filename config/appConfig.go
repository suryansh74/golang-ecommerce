package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerHost            string
	ServerPort            string
	Dsn                   string
	AppSecret             string
	TwilioAccountSID      string
	TwilioAuthToken       string
	TwilioFromPhoneNumber string
}

func SetupEnv() (cfg AppConfig, err error) {
	godotenv.Load()
	log.Println("printing all env vars")
	for _, env := range os.Environ() {
		log.Printf("%s\n", env)
	}

	// if os.Getenv("APP_ENV") == "dev" {
	// 	godotenv.Load()
	// }
	httpHost := os.Getenv("SERVER_HOST")
	httpPort := os.Getenv("SERVER_PORT")
	if len(httpHost) < 1 || len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variable are not set for host and port")
	}
	dsn := os.Getenv("DSN")
	appSecret := os.Getenv("APP_SECRET")

	return AppConfig{
		ServerHost:            httpHost,
		ServerPort:            httpPort,
		Dsn:                   dsn,
		AppSecret:             appSecret,
		TwilioAccountSID:      os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:       os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioFromPhoneNumber: os.Getenv("TWILIO_FROM_PHONE_NUMBER"),
	}, nil
}
