package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func DeleteResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}

func ResponseWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func NotFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"error": message})
}

func BadRequestResponse(c *gin.Context, message ...string) {

	if len(message) <= 0 {
		message = []string{"Bad Request"}
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": message[0]})
}

func ErrorResponse(c *gin.Context, message ...string) {
	if len(message) <= 0 {
		message = []string{"Internal Server Error"}
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": message[0]})
}
