package config

import (
	"fmt"
	"go-git-crud/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// Auto Migrate your models
	db.AutoMigrate(&models.Product{}) // Add your model structs here

	db.Create(&models.Product{Name: "Product 1", Price: 100})
	db.Create(&models.Product{Name: "Product 2", Price: 200})
	db.Create(&models.Product{Name: "Product 3", Price: 300})

	var products []models.Product
	db.Find(&products)
	fmt.Printf("products: %v\n", products)
	return db
}
