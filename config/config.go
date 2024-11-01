package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ApiKey string
}

func LoadConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading WEATHER_API_KEY")
	}

	apiKey := os.Getenv("WEATHER_API_KEY")

	return Config{ApiKey: apiKey}
}
