package middleware

import (
	"NonuNamak/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware проверяет токен и кладёт claims в контекст
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат токена"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
			c.Abort()
			return
		}

		// сохраняем данные о пользователе в контексте
		c.Set("userID", uint(claims.UserID))
		c.Set("role", claims.Role)

		c.Next()
	}

}

// AdminMiddleware разрешает доступ только админам
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещён: нужны права администратора"})
			c.Abort()
			return
		}
		c.Next()
	}
}
