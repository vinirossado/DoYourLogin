package middlewares

import (
	"context"
	"doYourLogin/source/domain/enumerations"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type ContextKey string

const ClaimsKey ContextKey = "jwtClaims"

var TokenClaims = &Claims{}

func ExtractClaims(c *gin.Context) *Claims {
	globalClaims := jwt.ExtractClaims(c)
	claims := Claims{
		ID:        uint(globalClaims["id"].(float64)),
		CompanyID: uint(globalClaims["company_id"].(float64)),
		Name:      globalClaims[identityKey].(string),
		Role:      enumerations.Roles(globalClaims["role"].(float64)),
	}

	ctx := context.WithValue(c.Request.Context(), ClaimsKey, &claims)
	c.Request = c.Request.WithContext(ctx)

	GetClaimsFromContext(c)

	return &claims
}

func GetClaimsFromContext(c *gin.Context) {
	TokenClaims, _ = c.Request.Context().Value(ClaimsKey).(*Claims)
}
