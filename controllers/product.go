package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pharmacy-backend/config"
	"pharmacy-backend/models"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении товаров"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки категорий"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func AddCategory(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil || input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Название обязательно"})
		return
	}

	category := models.Category{Name: input.Name}
	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения категории"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func SearchProducts(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр q обязателен"})
		return
	}

	var results []models.Product
	if err := config.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+query+"%").Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка поиска"})
		return
	}

	c.JSON(http.StatusOK, results)
}
