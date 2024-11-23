package controllers

import (
	"go-git-crud/middleware"
	"go-git-crud/models"
	"go-git-crud/services"
	"go-git-crud/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type MflixController struct {
	service *services.MflixService
}

func NewMflixController(db *mongo.Database) *MflixController {
	return &MflixController{service: services.NewMflixService(db)}
}

func (controller *MflixController) GetMflixs(c *gin.Context) {
	// Get pagination from context
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

	mflixs, err := controller.service.GetMflixs(paginationQuery)
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.Response(c, mflixs)
}

func (controller *MflixController) CreateMflix(c *gin.Context) {
	var mflix models.Mflix
	if err := c.ShouldBindJSON(&mflix); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}

	createdMflix, err := controller.service.CreateMflix(mflix)
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.Response(c, createdMflix)
}

func (controller *MflixController) UpdateMflix(c *gin.Context) {
	id := c.Param("id")
	var mflix models.Mflix
	if err := c.ShouldBindJSON(&mflix); err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}

	updatedMflix, err := controller.service.UpdateMflix(mflix, id)
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.Response(c, updatedMflix)
}

func (controller *MflixController) DeleteMflix(c *gin.Context) {
	id := c.Param("id")
	_, err := controller.service.DeleteMflix(id)
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}
	utils.DeleteResponse(c)
}
