package controllers

import (
	"fmt"
	"go-git-crud/services"
	"go-git-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}

	utils.Response(c, products)
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
