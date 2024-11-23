package config

import (
	"go-git-crud/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(GetEnv("SQLITE_DATABASE")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// Auto Migrate your models
	db.AutoMigrate(&models.Product{}) // Add your model structs here
	return db
}
