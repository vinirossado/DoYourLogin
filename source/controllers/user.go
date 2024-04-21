package controllers

import (
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/domain/responses"
	"doYourLogin/source/services"
	"doYourLogin/source/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	users := services.FindUsers()

	payload := responses.JsonResponse{
		Error:   false,
		Message: "Users found",
		Data:    users,
	}

	c.JSON(http.StatusOK, payload)
}

func FindUserById(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	user := services.FindUserById(id)

	payload := responses.JsonResponse{
		Error:   false,
		Message: "User found",
		Data:    user,
	}

	c.JSON(http.StatusOK, payload)
}

func CreateUser(c *gin.Context) {
	user := requests.UserRequest{}
	utils.ReadRequestBody(c, &user)

	services.CreateUser(&user)

	payload := responses.JsonResponse{
		Error:   false,
		Message: "User has been created",
	}

	c.JSON(http.StatusCreated, payload)
}

func UpdateUser(c *gin.Context) {
	updateUserRequest := requests.UserUpdateRequest{}
	utils.ReadRequestBody(c, &updateUserRequest)
	id := utils.ConvertToInt(c.Params.ByName("id"))

	services.UpdateUser(&updateUserRequest, id)

	payload := responses.JsonResponse{
		Error:   false,
		Message: "Users has been updated",
	}

	c.JSON(http.StatusOK, payload)
}

func DeleteUser(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	services.DeleteUser(id)

	payload := responses.JsonResponse{
		Error:   false,
		Message: "Users has been deleted",
	}

	c.JSON(http.StatusOK, payload)
}
