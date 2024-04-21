package routes

import (
	"doYourLogin/source/controllers"
	"doYourLogin/source/middlewares"
	"github.com/gin-gonic/gin"
)

func bindAuthorRoutes(router *gin.Engine) {

	users := router.Group("/author")

	users.Use(middlewares.JwtMiddleware().MiddlewareFunc())
	users.POST("", controllers.CreateAuthor)
	users.GET("", controllers.FindAuthor)
	users.GET("/:id", controllers.FindAuthorById)
	users.PUT("/:id", controllers.UpdateAuthor)
	users.PATCH("/:id", controllers.DeleteAuthor)
}
