package crypt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware для проверки токена
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get token from JSON Header
		tokenString := c.GetHeader("Authorization")
		//Check that token created and valid
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		// TODO send via protobuf to auth service
		// 	claims := jwt.MapClaims{}
		// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 		return secretKey, nil
		// 	})

		// 	if err != nil || !token.Valid {
		// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		// 		c.Abort()
		// 		return
		// 	}

		// 	c.Set("username", claims["username"])
		// 	c.Next()
	}
}
