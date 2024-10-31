package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	return port
}
