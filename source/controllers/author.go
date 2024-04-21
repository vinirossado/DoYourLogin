package controllers

import (
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/domain/responses"
	"doYourLogin/source/services"
	"doYourLogin/source/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAuthor(c *gin.Context) {
	authors := services.FindAuthors()

	payload := responses.JsonResponse{
		Error:   false,
		Message: "Authors found",
		Data:    authors,
	}

	c.JSON(http.StatusOK, payload)
}

func FindAuthorById(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	author := services.FindAuthorById(id)
	c.JSON(http.StatusOK, author)
}

func CreateAuthor(c *gin.Context) {
	author := requests.AuthorRequest{}
	utils.ReadRequestBody(c, &author)
	services.CreateAuthor(&author)

	payload := responses.JsonResponse{
		Error:   false,
		Message: "Author created",
		Data:    author,
	}

	c.JSON(http.StatusCreated, payload)
}

func UpdateAuthor(c *gin.Context) {
	updateUserRequest := requests.UserUpdateRequest{}
	utils.ReadRequestBody(c, &updateUserRequest)
	id := utils.ConvertToInt(c.Params.ByName("id"))

	services.UpdateUser(&updateUserRequest, id)
	c.Status(http.StatusOK)
}

func DeleteAuthor(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	services.DeleteUser(id)
	c.Status(http.StatusOK)
}
