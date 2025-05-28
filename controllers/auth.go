package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"pharmacy-backend/config"
	"pharmacy-backend/models"
	"pharmacy-backend/utils"
)

func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{Name: input.Name, Email: input.Email, Password: string(hashedPassword)}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь уже существует"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Регистрация успешна"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный пароль"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат"})
		return
	}

	if err := config.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Updates(models.User{Name: input.Name, Email: input.Email}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления профиля"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Профиль обновлен"})
}
