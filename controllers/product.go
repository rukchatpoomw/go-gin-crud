package controllers

import (
	"go-git-crud/middleware"
	"go-git-crud/models"
	"go-git-crud/services"
	"go-git-crud/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	service *services.ProductService
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{service: services.NewProductService(db)}
}

// Get all products
func (controller *ProductController) GetProducts(c *gin.Context) {
	pagination, exists := c.Get("pagination")
	if !exists {
		utils.BadRequestResponse(c, "page and limit must be greater than 0")
		return
	}

	// Make pagination has a struct
	paginationQuery, ok := pagination.(middleware.PaginationQuery)
	if !ok {
		utils.ErrorResponse(c, "invalid pagination")
		return
	}

	// fmt.Println("skip: ", paginationQuery.Limit)
	// fmt.Println("limit: ", limit)

	products, err := controller.service.GetProducts(paginationQuery)
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.Response(c, products)
}

// Get product by id
func (controller *ProductController) GetProduct(c *gin.Context) {
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
func (controller *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product

	// Bind JSON to product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.BadRequestResponse(c)
		return
	}

	result, err := controller.service.CreateProduct(product)
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}

	utils.Response(c, result)
}

// Update product
func (controller *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.BadRequestResponse(c)
		return
	}

	products, err := controller.service.UpdateProduct(product, id)
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
func (controller *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, err := controller.service.DeleteProduct(id)
	if err != nil {
		if err.Error() == "product not found" {
			utils.NotFoundResponse(c, err.Error())
		} else {
			utils.ErrorResponse(c, err.Error())
		}
		return
	}
	utils.DeleteResponse(c)
}
