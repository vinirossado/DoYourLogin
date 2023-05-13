package controllers

import (
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/services"
	"doYourLogin/source/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCompany(c *gin.Context) {
	company := requests.CompanyRequest{}
	utils.ReadRequestBody(c, &company)
	company2 := services.CreateCompany(&company)
	c.JSON(http.StatusOK, company2)
}
