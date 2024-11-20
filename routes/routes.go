package routes

import (
	"fmt"
	"go-git-crud/config"
	"go-git-crud/controllers"
	"go-git-crud/repositories"
	"go-git-crud/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	config.LoadEnv()
	mongodb := config.ConnectMongoDB()
	fmt.Printf("mongodb: %v\n", mongodb)
	sql := config.ConnectSQLite()

	productRepo := repositories.ProductRepository(sql)
	productService := services.ProductService(productRepo)
	productController := controllers.ProductController(productService)

	r.GET("/products", productController.GetProducts)
	r.GET("/product/:id", productController.GetProduct)
	r.POST("/product", productController.CreateProduct)
	r.PUT("/product/:id", controllers.UpdateProduct)
	r.DELETE("/product/:id", controllers.DeleteProduct)
}
