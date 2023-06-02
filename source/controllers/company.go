package controllers

import (
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/domain/responses"
	"doYourLogin/source/services"
	"doYourLogin/source/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCompany(c *gin.Context) {
	company := requests.CompanyRequest{}
	utils.ReadRequestBody(c, &company)
	companyResponse := services.CreateCompany(&company)
	c.JSON(http.StatusOK, companyResponse)
}

func FindCompanies(c *gin.Context) {
	companies := services.FindCompanies()
	response := responses.Response{
		StatusCode: http.StatusOK,
		Data:       companies,
	}
	c.JSON(http.StatusOK, response)
}

func FindMyCompany(c *gin.Context) {
	company := services.FindMyCompany()
	response := responses.Response{
		StatusCode: http.StatusOK,
		Data:       company,
	}
	c.JSON(http.StatusOK, response)
}

func ActivateAccount(c *gin.Context) {
	tokenAPi := utils.ReadBodyParam(c, "token")
	services.ActivateAccount(tokenAPi)
	c.JSON(http.StatusOK, nil)
}
