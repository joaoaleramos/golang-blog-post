package utils

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type DBConfig struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
	DbSSLMode  string
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		DbUser:     os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		DbSSLMode:  os.Getenv("DB_SSL_MODE"),
	}
}
