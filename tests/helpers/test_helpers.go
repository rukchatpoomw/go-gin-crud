package helpers

import (
	"go-git-crud/controllers"
	"go-git-crud/middleware"
	"go-git-crud/models"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupTestDB creates and configures an in-memory SQLite database for testing
func SetupTestDB(t *testing.T) *gorm.DB {
	// Use in-memory SQLite for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	// Create test data
	testProduct := models.Product{
		Name:  "Test Product",
		Price: 100,
	}
	if err := db.Create(&testProduct).Error; err != nil {
		t.Fatalf("failed to create test product: %v", err)
	}

	return db
}

// SetupTestRouter creates and configures a Gin router for testing
func SetupTestRouter(controller *controllers.ProductController) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Recovery())

	// Product routes with pagination middleware for GET /products
	r.GET("/products", middleware.PaginationMiddleware(), controller.GetProducts)
	r.GET("/products/:id", controller.GetProduct)
	r.POST("/products", controller.CreateProduct)
	r.PUT("/products/:id", controller.UpdateProduct)
	r.DELETE("/products/:id", controller.DeleteProduct)

	return r
}
