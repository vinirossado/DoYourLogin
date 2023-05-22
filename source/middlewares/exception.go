package middlewares

import (
	"doYourLogin/source/domain/exceptions"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   any    `json:"error"`
}

func ExceptionMiddleware(c *gin.Context, recovered interface{}) {
	if except, ok := recovered.(*exceptions.HttpException); ok {
		errorResponse := ErrorResponse{
			Message: except.Message,
			Code:    except.StatusCode,
			Error:   "Something went wrong during processing, try again later",
		}
		c.JSON(except.StatusCode, errorResponse)
	} else {
		log.Printf("Exception not mapped: %s", recovered)
		errorResponse := ErrorResponse{
			Message: except.Message,
			Code:    except.StatusCode,
			Error:   recovered,
		}
		c.JSON(http.StatusInternalServerError, errorResponse)
	}
}
