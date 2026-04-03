package main

import (
	"marketplace_backend/internal/database"
	"marketplace_backend/internal/models"
	"marketplace_backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	database.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
	)

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
