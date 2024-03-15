package database

import (
	"assignment-2/config"
	"assignment-2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() (Database, error) {
	dbInfo := config.GetDatabaseEnv()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbInfo.Host, dbInfo.Port, dbInfo.User, dbInfo.Password, dbInfo.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return Database{}, err
	}

	if err := db.Debug().AutoMigrate(&models.Order{}, &models.Item{}); err != nil {
		fmt.Println("Error performing database migrations:", err)
		return Database{}, err
	}

	return Database{db: db}, nil
}
