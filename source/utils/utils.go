package utils

import (
	"abrigos/source/domain/exception"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadRequestBody(c *gin.Context, requestBody interface{}) {

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		exception.ThrowBadRequestException(err.Error())
	}

}

func ConvertToInt(stringValue string) int {
	valueConv, err := strconv.Atoi(stringValue)

	if err != nil {
		exception.ThrowBadRequestException(fmt.Sprintf("Error converting parameter to int with error: %s", err))
	}

	return valueConv
}
