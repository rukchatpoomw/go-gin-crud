package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-git-crud/controllers"
	"go-git-crud/middleware"
	"go-git-crud/models"
	"go-git-crud/tests/cases"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
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

func setupTestRouter(controller *controllers.ProductController) *gin.Engine {
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

// TestGetProducts tests the GET /products endpoint with different scenarios
func TestGetProducts(t *testing.T) {
	// Define test cases using table-driven test pattern
	tests := cases.GetProductsTestCases()

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// Setup test environment
			db := setupTestDB(t)                               // Create in-memory SQLite database
			controller := controllers.NewProductController(db) // Create controller instance
			router := setupTestRouter(controller)              // Setup Gin router with routes

			// Create HTTP test recorder and request
			w := httptest.NewRecorder() // Mock HTTP response writer
			// Create GET request with pagination query parameters
			req := httptest.NewRequest(http.MethodGet, "/products?page="+tt.Page+"&limit="+tt.Limit, nil)
			router.ServeHTTP(w, req) // Process the request

			// Verify response status code
			if w.Code != tt.ExpectedCode {
				t.Errorf("Expected status code %d, got %d", tt.ExpectedCode, w.Code)
			}

			// For successful requests, verify response body
			if w.Code == http.StatusOK {
				var response []models.Product
				// Parse JSON response into Product slice
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				// Verify number of products returned
				if len(response) != tt.ExpectedCount {
					t.Errorf("Expected %d products, got %d", tt.ExpectedCount, len(response))
				}

				// Verify product data types and constraints
				for _, product := range response {
					// Check ID is not zero
					if product.ID == 0 {
						t.Error("Product ID should not be zero")
					}

					// Check Name is not empty
					if product.Name == "" {
						t.Error("Product Name should not be empty")
					}

					// Check Price is positive
					if product.Price <= 0 {
						t.Errorf("Product Price should be positive, got %f", product.Price)
					}

					// Check CreatedAt is not zero
					if product.CreatedAt.IsZero() {
						t.Error("Product CreatedAt should not be zero")
					}

					// Check UpdatedAt is not zero
					if product.UpdatedAt.IsZero() {
						t.Error("Product UpdatedAt should not be zero")
					}
				}
			} else {
				// For error responses, verify error message format
				var errorResponse map[string]interface{}
				if err := json.Unmarshal(w.Body.Bytes(), &errorResponse); err == nil {
					// Check if error message exists
					if errorMsg, exists := errorResponse["error"]; !exists {
						t.Error("Error response should contain 'error' field")
					} else {
						// Check if error message is string
						if _, ok := errorMsg.(string); !ok {
							t.Error("Error message should be a string")
						}
					}
				}
			}
		})
	}
}

// func TestGetProduct(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		productID    string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Success",
// 			productID:    "1",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Not Found",
// 			productID:    "999",
// 			expectedCode: http.StatusNotFound,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			db := setupTestDB(t)
// 			controller := NewProductController(db)
// 			router := setupTestRouter(controller)

// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodGet, "/products/"+tt.productID, nil)
// 			router.ServeHTTP(w, req)

// 			if w.Code != tt.expectedCode {
// 				t.Errorf("Expected status code %d, got %d", tt.expectedCode, w.Code)
// 			}

// 			if w.Code == http.StatusOK {
// 				var response models.Product
// 				err := json.Unmarshal(w.Body.Bytes(), &response)
// 				if err != nil {
// 					t.Fatalf("Failed to unmarshal response: %v", err)
// 				}

// 				if response.Name != "Test Product" {
// 					t.Errorf("Expected product name 'Test Product', got '%s'", response.Name)
// 				}
// 			}
// 		})
// 	}
// }

// func TestCreateProduct(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		product      models.Product
// 		expectedCode int
// 	}{
// 		{
// 			name: "Success",
// 			product: models.Product{
// 				Name:  "New Product",
// 				Price: 200,
// 			},
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Invalid Request - Empty Product",
// 			product:      models.Product{},
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name: "Invalid Request - Missing Price",
// 			product: models.Product{
// 				Name: "Invalid Product",
// 			},
// 			expectedCode: http.StatusBadRequest,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			db := setupTestDB(t)
// 			controller := NewProductController(db)
// 			router := setupTestRouter(controller)

// 			body, _ := json.Marshal(tt.product)
// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
// 			req.Header.Set("Content-Type", "application/json")
// 			router.ServeHTTP(w, req)

// 			if w.Code != tt.expectedCode {
// 				t.Errorf("Expected status code %d, got %d", tt.expectedCode, w.Code)
// 			}

// 			if w.Code == http.StatusOK {
// 				var response models.Product
// 				err := json.Unmarshal(w.Body.Bytes(), &response)
// 				if err != nil {
// 					t.Fatalf("Failed to unmarshal response: %v", err)
// 				}

// 				if response.Name != tt.product.Name {
// 					t.Errorf("Expected product name '%s', got '%s'", tt.product.Name, response.Name)
// 				}
// 			}
// 		})
// 	}
// }

// func TestUpdateProduct(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		productID    string
// 		product      models.Product
// 		setupDB      func(*gorm.DB)
// 		expectedCode int
// 	}{
// 		{
// 			name:      "Success",
// 			productID: "1",
// 			product: models.Product{
// 				Name:  "Updated Product",
// 				Price: 300,
// 			},
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:      "Not Found",
// 			productID: "999",
// 			product: models.Product{
// 				Name:  "Updated Product",
// 				Price: 300,
// 			},
// 			expectedCode: http.StatusNotFound,
// 			setupDB: func(db *gorm.DB) {
// 				// Delete all products to ensure 404
// 				db.Exec("DELETE FROM products")
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			db := setupTestDB(t)
// 			if tt.setupDB != nil {
// 				tt.setupDB(db)
// 			}

// 			controller := NewProductController(db)
// 			router := setupTestRouter(controller)

// 			body, _ := json.Marshal(tt.product)
// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodPut, "/products/"+tt.productID, bytes.NewBuffer(body))
// 			req.Header.Set("Content-Type", "application/json")
// 			router.ServeHTTP(w, req)

// 			if w.Code != tt.expectedCode {
// 				t.Errorf("Expected status code %d, got %d", tt.expectedCode, w.Code)
// 			}

// 			if w.Code == http.StatusOK {
// 				var response models.Product
// 				err := json.Unmarshal(w.Body.Bytes(), &response)
// 				if err != nil {
// 					t.Fatalf("Failed to unmarshal response: %v", err)
// 				}

// 				if response.Name != tt.product.Name {
// 					t.Errorf("Expected product name '%s', got '%s'", tt.product.Name, response.Name)
// 				}
// 			}
// 		})
// 	}
// }

// func TestDeleteProduct(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		productID    string
// 		setupDB      func(*gorm.DB)
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Success",
// 			productID:    "1",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Not Found",
// 			productID:    "999",
// 			expectedCode: http.StatusNotFound,
// 			setupDB: func(db *gorm.DB) {
// 				// Delete all products to ensure 404
// 				db.Exec("DELETE FROM products")
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			db := setupTestDB(t)
// 			if tt.setupDB != nil {
// 				tt.setupDB(db)
// 			}

// 			controller := NewProductController(db)
// 			router := setupTestRouter(controller)

// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodDelete, "/products/"+tt.productID, nil)
// 			router.ServeHTTP(w, req)

// 			if w.Code != tt.expectedCode {
// 				t.Errorf("Expected status code %d, got %d", tt.expectedCode, w.Code)
// 			}
// 		})
// 	}
// }
