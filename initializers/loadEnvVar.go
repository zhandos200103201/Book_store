package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVar() {
	if err := godotenv.Load("initializers/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
