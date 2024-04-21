package controllers

import (
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/domain/responses"
	"doYourLogin/source/services"
	"doYourLogin/source/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindBooks(c *gin.Context) {
	books := services.FindBooks()

	payload := responses.JsonResponse{
		Error:   false,
		Message: "Books found",
		Data:    books,
	}

	c.JSON(http.StatusOK, payload)
}

func FindBookById(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	book := services.FindBookById(id)

	payload := responses.JsonResponse{
		Error:   false,
		Message: "Book found",
		Data:    book,
	}

	c.JSON(http.StatusOK, payload)
}

func CreateBook(c *gin.Context) {
	book := requests.BookRequest{}
	utils.ReadRequestBody(c, &book)

	services.CreateBook(&book)

	payload := responses.JsonResponse{
		Error:   false,
		Message: "Book has been created",
		Data:    book,
	}

	c.JSON(http.StatusCreated, payload)
}

func UpdateBook(c *gin.Context) {
	updateUserRequest := requests.UserUpdateRequest{}
	utils.ReadRequestBody(c, &updateUserRequest)
	id := utils.ConvertToInt(c.Params.ByName("id"))

	services.UpdateUser(&updateUserRequest, id)

	c.Status(http.StatusOK)
}

func DeleteBook(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	services.DeleteUser(id)

	c.Status(http.StatusOK)
}
