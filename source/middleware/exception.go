package middleware

import (
	"abrigos/source/domain/exception"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExceptionMiddleware(c *gin.Context, recovered interface{}) {
	if except, ok := recovered.(*exception.HttpException); ok {
		c.String(except.StatusCode, except.Message)
	} else {
		log.Printf("Exception not mapped: %s", recovered)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
