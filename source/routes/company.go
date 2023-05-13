package routes

import (
	"doYourLogin/source/controllers"
	"github.com/gin-gonic/gin"
)

func bindCompanyRoutes(router *gin.Engine) {

	company := router.Group("/company")

	company.POST("", controllers.CreateCompany)

}
