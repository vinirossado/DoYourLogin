package routes

import (
	"doYourLogin/source/controllers"
	"doYourLogin/source/domain/enumerations"
	"doYourLogin/source/middlewares"
	"github.com/gin-gonic/gin"
)

func bindCompanyRoutes(router *gin.Engine) {

	company := router.Group("/company")

	company.POST("", controllers.CreateCompany)

	company.Use(middlewares.JwtMiddleware().MiddlewareFunc())
	company.GET("/my-company", middlewares.AuthorizationMiddleware(enumerations.ADMIN), controllers.FindMyCompany)

	company.GET("", middlewares.AuthorizationMiddleware(enumerations.GOD), controllers.FindCompanies)

}
