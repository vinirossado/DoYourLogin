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
	companyResponse := services.CreateCompany(&company)
	c.JSON(http.StatusOK, companyResponse)
}

func FindCompanies(c *gin.Context) {
	companies := services.FindCompanies()
	c.JSON(http.StatusOK, companies)
}

func FindCompanyByID(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	company := services.FindCompanyByID(id)
	c.JSON(http.StatusOK, company)
}
