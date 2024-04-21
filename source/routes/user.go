package routes

import (
	"doYourLogin/source/controllers"
	"doYourLogin/source/middlewares"

	"github.com/gin-gonic/gin"
)

func bindUserRoutes(router *gin.Engine) {

	users := router.Group("/user")

	users.POST("", controllers.CreateUser)
	users.Use(middlewares.JwtMiddleware().MiddlewareFunc())
	users.GET("", controllers.FindUsers)
	users.GET("/:id", controllers.FindUserById)
	users.PUT("/:id", controllers.UpdateUser)
	users.PATCH("/:id", controllers.DeleteUser)
}
