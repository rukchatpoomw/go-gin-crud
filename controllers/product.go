package controllers

import (
	"fmt"
	"go-git-crud/models"
	"go-git-crud/services"
	"go-git-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all products
func GetProducts(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}

	utils.Response(c, products)
}

// Get product by id
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := services.GetProduct(id)

	// Handle error if product not found
	// if err.Error() == "product not found" {
	// 	utils.NotFoundResponse(c, err.Error())
	// 	return
	// }

	// Handle error if not found item then any error should be handled

	if err != nil {
		if err.Error() == "product not found" {
			utils.NotFoundResponse(c, err.Error())
		} else {
			utils.ErrorResponse(c, err.Error())
		}
		return
	}

	utils.Response(c, product)
}

// Create product
func CreateProduct(c *gin.Context) {
	var product models.Product
	fmt.Println("product:", product)

	// Bind JSON to product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	fmt.Println("product:", product)

	product, err := services.CreateProduct(product)
	if err != nil {
		if err.Error() == "product with the same ID already exists" {
			utils.NotFoundResponse(c, err.Error())
		} else {
			utils.ErrorResponse(c, err.Error())
		}
		return
	}
	utils.Response(c, product)
}

// Update product
func UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Product",
	})
}

// Delete product
func DeleteProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Product",
	})
}
