package routes

import (
	"marketplace_backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUser)
	r.GET("/users", handlers.GetUsers)

	r.POST("/products", handlers.CreateProduct)
	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	r.POST("/categories", handlers.CreateCategory)
	r.GET("/categories", handlers.GetCategories)

}
