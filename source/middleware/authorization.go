package middleware

import (
	"abrigos/source/domain/enumerations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(minimumRole enumerations.Roles) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := ExtractClaims(c)

		if claims.Role < minimumRole {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
