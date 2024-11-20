package controllers

import (
	"go-git-crud/models"
	"go-git-crud/services"
	"go-git-crud/utils"

	"github.com/gin-gonic/gin"
)

type ProductControllerType struct {
	service *services.ProductServiceType
}

func ProductController(service *services.ProductServiceType) *ProductControllerType {
	return &ProductControllerType{service: service}
}

// Get all products
func (controller *ProductControllerType) GetProducts(c *gin.Context) {
	products, err := controller.service.GetProducts()
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.Response(c, products)
}

// Get product by id
func (controller *ProductControllerType) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := controller.service.GetProduct(id)

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
func (controller *ProductControllerType) CreateProduct(c *gin.Context) {
	var product models.Product

	// Bind JSON to product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.BadRequestResponse(c)
		return
	}

	// utils.Response(c, result)
}

// Update product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.BadRequestResponse(c)
		return
	}

	products, err := services.UpdateProduct(product, id)
	if err != nil {
		if err.Error() == "product not found" {
			utils.NotFoundResponse(c, err.Error())
		} else {
			utils.ErrorResponse(c, err.Error())
		}
		return
	}
	utils.Response(c, products)
}

// Delete product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	products, err := services.DeleteProduct(id)
	if err != nil {
		if err.Error() == "product not found" {
			utils.NotFoundResponse(c, err.Error())
		} else {
			utils.ErrorResponse(c, err.Error())
		}
		return
	}
	utils.Response(c, products)
}
