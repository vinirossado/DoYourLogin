package routes

import (
	"abrigos/source/controllers"
	"abrigos/source/domain/enumerations"
	"abrigos/source/middleware"

	"github.com/gin-gonic/gin"
)

func bindUserRoutes(router *gin.Engine) {

	users := router.Group("/tutors")

	users.POST("", controllers.CreateUser)

	users.Use(middleware.JwtMiddleware().MiddlewareFunc())
	users.GET("", middleware.AuthorizationMiddleware(enumerations.NORMAL), controllers.FindUsers)
	users.GET("/:id", middleware.AuthorizationMiddleware(enumerations.NORMAL), controllers.FindUserById)
	users.PUT("/:id", middleware.AuthorizationMiddleware(enumerations.NORMAL), controllers.UpdateUser)
	users.PATCH("/:id", middleware.AuthorizationMiddleware(enumerations.NORMAL), controllers.DeleteUser)
}
