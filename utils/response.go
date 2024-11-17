package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func ResponseWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func NotFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"error": message})
}

func ErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": message})
}
