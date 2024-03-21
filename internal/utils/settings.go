package utils

import (
	"os"

	"log"

	"github.com/joho/godotenv"
)

type Settings struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
	DbSSLMode  string
}

func NewSettings() *Settings {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Settings{
		DbUser:     os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		DbSSLMode:  os.Getenv("DB_SSL_MODE"),
	}
}
