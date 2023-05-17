package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiToken := c.GetHeader("ApiToken")
		if apiToken == "" {
			authorizationHeader := c.GetHeader("Authorization")
			if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "ApiToken or BearerToken is required"})
				c.Abort()
				return
			}
		}

		// You can access the value of the ApiToken or bearer token and use it as needed in your code.

		// Continue to the next middleware or route handler
		c.Next()
	}
}
