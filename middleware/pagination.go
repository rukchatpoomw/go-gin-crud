package middleware

import (
	"go-git-crud/utils"

	"github.com/gin-gonic/gin"
)

type PaginationQuery struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	Skip  int64 `json:"skip"`
}

func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get query parameters with defaults
		page := c.DefaultQuery("page", "1")
		limit := c.DefaultQuery("limit", "10")

		// Convert to int64
		pageInt64 := utils.ConvertStringToInt64(page, 1)
		limitInt64 := utils.ConvertStringToInt64(limit, 10)

		// Handle invalid values
		// If invalid break middleware
		if pageInt64 <= 0 || limitInt64 <= 0 {
			return
		}

		// Calculate skip
		skip := utils.CalculateSkip(pageInt64, limitInt64)

		// Set pagination in context
		c.Set("pagination", PaginationQuery{
			Page:  pageInt64,
			Limit: limitInt64,
			Skip:  skip,
		})

		c.Next()
	}
}
