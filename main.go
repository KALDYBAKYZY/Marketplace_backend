package main

import (
	"marketplace_backend/internal/database"
	"marketplace_backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
