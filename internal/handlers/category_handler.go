package handlers

import (
	"marketplace_backend/internal/database"
	"marketplace_backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var existing models.Category
	if err := database.DB.Where("name = ?", category.Name).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Category with this name already exists"})
		return
	}
	database.DB.Create(&category)
	c.JSON(http.StatusCreated, category)
}

func GetCategories(c *gin.Context) {
	var categories []models.Category
	database.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}
