package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DatabaseInfo struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}
}

func GetServerPortEnv() string {
	LoadEnvFile()
	return os.Getenv("SERVER_PORT")
}

func GetDatabaseEnv() DatabaseInfo {
	LoadEnvFile()
	return DatabaseInfo{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}
