package routes

import (
	"doYourLogin/source/controllers"
	"doYourLogin/source/domain/enumerations"
	"doYourLogin/source/middlewares"

	"github.com/gin-gonic/gin"
)

func bindUserRoutes(router *gin.Engine) {

	users := router.Group("/user")

	users.Use(middlewares.JwtMiddleware().MiddlewareFunc())
	users.POST("", middlewares.AuthorizationMiddleware(enumerations.ADMIN), controllers.CreateUser)
	users.GET("", middlewares.AuthorizationMiddleware(enumerations.NORMAL), controllers.FindUsers)
	users.GET("/:id", middlewares.AuthorizationMiddleware(enumerations.NORMAL), controllers.FindUserById)
	users.PUT("/:id", middlewares.AuthorizationMiddleware(enumerations.NORMAL), controllers.UpdateUser)
	users.PATCH("/:id", middlewares.AuthorizationMiddleware(enumerations.ADMIN), controllers.DeleteUser)
}
