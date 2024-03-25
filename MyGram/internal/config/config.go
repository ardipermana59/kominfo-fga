package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	return nil
}
