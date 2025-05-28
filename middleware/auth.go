package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret") // Тот же ключ, что и в utils/token.go

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получение токена из заголовка
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не найден"})
			c.Abort()
			return
		}

		// Извлечение токена из "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат токена"})
			c.Abort()
			return
		}

		tokenStr := parts[1]

		// Разбор и верификация токена
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недопустимый токен"})
			c.Abort()
			return
		}

		// Извлекаем user_id из claims и сохраняем в контексте
		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка авторизации"})
			c.Abort()
			return
		}

		c.Set("user_id", uint(userID))
		c.Next()
	}
}
