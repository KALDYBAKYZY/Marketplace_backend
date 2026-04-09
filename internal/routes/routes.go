package routes

import (
	"marketplace_backend/internal/handlers"
	"marketplace_backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	{
		protected.POST("/users", handlers.CreateUser)
		protected.GET("/users/:id", handlers.GetUser)
		protected.GET("/users", handlers.GetUsers)

		protected.POST("/products", handlers.CreateProduct)
		protected.GET("/products", handlers.GetProducts)
		protected.GET("/products/:id", handlers.GetProduct)
		protected.PUT("/products/:id", handlers.UpdateProduct)
		protected.DELETE("/products/:id", handlers.DeleteProduct)

		protected.POST("/categories", handlers.CreateCategory)
		protected.GET("/categories", handlers.GetCategories)
	}
}
