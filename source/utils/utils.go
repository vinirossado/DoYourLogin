package utils

import (
	"doYourLogin/source/domain/exceptions"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadRequestBody(c *gin.Context, requestBody interface{}) {
	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		exceptions.ThrowBadRequestException(err.Error())
	}
}

func ConvertToInt(stringValue string) int {
	valueConv, err := strconv.Atoi(stringValue)

	if err != nil {
		exceptions.ThrowBadRequestException(fmt.Sprintf("Error converting parameter to int with error: %s", err))
	}

	return valueConv
}

func ReadBodyParam(c *gin.Context, param string) string {
	return c.Params.ByName("token")
}
