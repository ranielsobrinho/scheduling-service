package router

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Token not found"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid Authorization header format"})
			return
		}
		tokenString := parts[1]

		jwtSecret := os.Getenv("JWT_SECRET_KEY")
		if jwtSecret == "" {
			c.AbortWithStatusJSON(500, gin.H{"message": "JWT secret key not configured"})
			return
		}

		payload, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !payload.Valid {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid token"})
			return
		}

		c.Next()
	}
}
