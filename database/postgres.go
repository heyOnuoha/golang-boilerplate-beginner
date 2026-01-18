package database

import (
	"errors"
	"fmt"
	"todo-api/config"
	"todo-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var AllModels = []interface{}{
	&models.TodoItem{},
	&models.TodoNote{},
	&models.User{},
}

func InitDatabase(config *config.DatabaseConfig) error {
	var err error
	// Use := only for the first declaration, not for assigning to the global DB variable
	DB, err = gorm.Open(postgres.Open(config.GetDatabaseString()), &gorm.Config{})

	if err != nil {
		return err
	}

	if posgresDB, err := DB.DB(); err == nil {
		posgresDB.SetMaxIdleConns(10)
		posgresDB.SetMaxOpenConns(100)

		return err
	}

	return nil
}

func Migrate() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}

	for _, model := range AllModels {
		if err := DB.AutoMigrate(model); err != nil {
			fmt.Printf("Failed to migrate model: %T, error: %v\n", model, err)
			return err
		}
	}

	return nil
}

func CloseDB() {
	if DB != nil {
		posgresDB, error := DB.DB()
		if error != nil {
			return
		}
		posgresDB.Close()
	}
}

func GetDB() *gorm.DB {
	return DB
}
