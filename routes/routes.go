package routes

import (
	"go-git-crud/config"
	"go-git-crud/controllers"
	"go-git-crud/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	config.LoadEnv()
	mongodb := config.ConnectMongoDB()
	sqlite := config.ConnectSQLite()

	// Product routes (SQLite)
	productController := controllers.NewProductController(sqlite)

	productRoutes := r.Group("/products")
	{
		productRoutes.GET("", middleware.PaginationMiddleware(), productController.GetProducts)
		productRoutes.GET("/:id", productController.GetProduct)
		productRoutes.POST("", productController.CreateProduct)
		productRoutes.PUT("/:id", productController.UpdateProduct)
		productRoutes.DELETE("/:id", productController.DeleteProduct)
	}

	mflixController := controllers.NewMflixController(mongodb)

	mflixRoutes := r.Group("/comments")
	{
		mflixRoutes.GET("", middleware.PaginationMiddleware(), mflixController.GetMflixs)
		mflixRoutes.POST("", mflixController.CreateMflix)
		mflixRoutes.PUT("/:id", mflixController.UpdateMflix)
		mflixRoutes.DELETE("/:id", mflixController.DeleteMflix)
	}
}
