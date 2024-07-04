package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saffrondigits/api/security"
)

func TokenMiddleware() gin.HandlerFunc {
	return Middleware
}

func Middleware(c *gin.Context) {
	// get Authorization token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	if err := security.VerifyJWTToken(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Next()
}

// In Golang you can pass a function/method as a parameter to another function/method
// In Golang you can return a function/method from a function/method
