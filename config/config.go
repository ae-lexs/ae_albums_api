package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	DBHost      string
	DBName      string
	DBPassword  string
	DBPort      string
	DBUser      string
	Environment string
}

func GetConfig() config {
	environment := os.Getenv("ENVIRONMENT")

	if environment == "" || environment == "DEVELOPMENT" {
		godotenv.Load()
	}

	return config{
		DBHost:      os.Getenv("DB_HOST"),
		DBName:      os.Getenv("DB_NAME"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		Environment: environment,
	}
}
