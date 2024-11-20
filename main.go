package main

import (
	"go-git-crud/config"
	routes "go-git-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	port := config.GetEnv("PORT", "8000")
	r.Run(":" + port)
}
