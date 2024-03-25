package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetDBConfig() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := map[string]string{
		"host":     os.Getenv("DB_HOST"),
		"user":     os.Getenv("DB_USER"),
		"password": os.Getenv("DB_PASSWORD"),
		"port":     os.Getenv("DB_PORT"),
		"dbname":   os.Getenv("DB_NAME"),
		"sslmode":  os.Getenv("DB_SSLMODE"),
	}

	connStr := make([]string, 0, len(config))
	for k, v := range config {
		connStr = append(connStr, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(connStr, " ")
}
