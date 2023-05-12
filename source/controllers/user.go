package controllers

import (
	"abrigos/source/domain/request"
	"abrigos/source/service"
	"abrigos/source/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	users := service.FindUsers()
	c.JSON(http.StatusOK, users)
}

func FindUserById(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	user := service.FindUserById(id)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	user := request.UserRequest{}
	utils.ReadRequestBody(c, &user)
	service.CreateUser(&user)
	c.Status(http.StatusOK)
}

func UpdateUser(c *gin.Context) {
	updateUserRequest := request.UserRequest{}
	utils.ReadRequestBody(c, &updateUserRequest)

	id := utils.ConvertToInt(c.Params.ByName("id"))
	service.UpdateUser(&updateUserRequest, id)
	c.Status(http.StatusOK)
}

func DeleteUser(c *gin.Context) {
	id := utils.ConvertToInt(c.Params.ByName("id"))
	service.DeleteUser(id)
	c.Status(http.StatusOK)
}
