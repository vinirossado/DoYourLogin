package controllers

import (
	"doYourLogin/source/domain/entities"
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/services"
	"doYourLogin/source/utils"
	"fmt"
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

	user := entities.User{
		Name:     "AAAAAAA",
		Username: "user1",
		Email:    "dycjh@example.com",
	}

	dto := &requests.UserRequest{}

	err := utils.Map(user, dto)

	if err != nil {
		fmt.Println("Mapping error:", err)
	}

	fmt.Println(dto)

	users := []entities.User{}
	dtos := []entities.User{
		{
			Name:     "user1",
			Username: "user1",
			Email:    "dycjh@example.com",
		},
		{
			Name:     "user1",
			Username: "user1",
			Email:    "dycjh@example.com",
		},
		{
			Name:     "user1",
			Username: "user1",
			Email:    "dycjh@example.com",
		},
	}

	err = utils.Map(users, dtos)
	if err != nil {
		fmt.Println("Mapping error:", err)
	}

	fmt.Print(dtos)

	companies := services.FindCompanies()
	c.JSON(http.StatusOK, companies)
}

func FindCompanyByID(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	company := services.FindCompanyByID(id)
	c.JSON(http.StatusOK, company)
}
