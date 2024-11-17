package controllers

import (
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
