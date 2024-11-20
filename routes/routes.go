package routes

import (
	"go-git-crud/config"
	"go-git-crud/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	config.LoadEnv()
	db := config.ConnectDB()
	productController := controllers.ProductController(db)

	r.GET("/products", productController.GetProducts)
	r.GET("/product/:id", controllers.GetProduct)
	r.POST("/product", controllers.CreateProduct)
	r.PUT("/product/:id", controllers.UpdateProduct)
	r.DELETE("/product/:id", controllers.DeleteProduct)
}
