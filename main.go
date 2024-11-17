package main

import (
	handlers "go-git-crud/handles"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/products", handlers.GetProducts)
	router.GET("/product/:id", handlers.GetProduct)
	router.POST("/product", handlers.CreateProduct)
	router.PUT("/product/:id", handlers.UpdateProduct)
	router.DELETE("/product/:id", handlers.DeleteProduct)

	router.Run()
}
