package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pharmacy-backend/config"
	"pharmacy-backend/models"
)

func CreateOrder(c *gin.Context) {
	var input struct {
		UserID  uint   `json:"user_id"`
		Address string `json:"address"`
		Phone   string `json:"phone"`
		Items   []struct {
			ProductID uint `json:"product_id"`
			Quantity  int  `json:"quantity"`
		}
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	order := models.Order{UserID: input.UserID, Address: input.Address, Phone: input.Phone}
	config.DB.Create(&order)

	for _, item := range input.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		config.DB.Create(&orderItem)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Заказ оформлен"})
}

func GetUserOrders(c *gin.Context) {
	var input struct {
		UserID uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат"})
		return
	}

	var orders []models.Order
	if err := config.DB.Preload("OrderItems").Where("user_id = ?", input.UserID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении заказов"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
