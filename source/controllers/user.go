package controllers

import (
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/services"
	"doYourLogin/source/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	users := services.FindUsers()
	c.JSON(http.StatusOK, users)
}

func FindUserById(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	user := services.FindUserById(id)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	user := requests.UserRequest{}
	utils.ReadRequestBody(c, &user)
	services.CreateUser(&user)
	c.Status(http.StatusOK)
}

func UpdateUser(c *gin.Context) {
	updateUserRequest := requests.UserUpdateRequest{}
	utils.ReadRequestBody(c, &updateUserRequest)
	id := utils.ConvertToInt(c.Params.ByName("id"))
	services.UpdateUser(&updateUserRequest, id)
	c.Status(http.StatusOK)
}

func DeleteUser(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	services.DeleteUser(id)
	c.Status(http.StatusOK)
}
