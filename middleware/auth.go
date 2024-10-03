package middleware

import (
	"Image-Processing-Service/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	// check auth header format
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	//verify auth token
	err := services.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired JWT token"})
		c.Abort()
		return
	}

	//continue to remaining handler
	c.Next()

}
