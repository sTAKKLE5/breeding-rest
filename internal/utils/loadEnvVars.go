package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVars() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
