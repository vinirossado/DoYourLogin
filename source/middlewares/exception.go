package middlewares

import (
	"doYourLogin/source/domain/exceptions"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExceptionMiddleware(c *gin.Context, recovered interface{}) {
	if except, ok := recovered.(*exceptions.HttpException); ok {
		c.String(except.StatusCode, except.Message)
	} else {
		log.Printf("Exception not mapped: %s", recovered)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
