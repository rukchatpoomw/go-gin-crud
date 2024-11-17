package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Products",
	})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Get Product %v", id),
	})
}

func CreateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Product",
	})
}

func UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Product",
	})
}

func DeleteProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Product",
	})
}
