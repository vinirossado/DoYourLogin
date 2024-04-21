package routes

import (
	"doYourLogin/source/controllers"
	"doYourLogin/source/middlewares"
	"github.com/gin-gonic/gin"
)

func bindBookRoutes(router *gin.Engine) {

	users := router.Group("/book")

	users.Use(middlewares.JwtMiddleware().MiddlewareFunc())
	users.POST("", controllers.CreateBook)
	users.GET("", controllers.FindBooks)
	users.GET("/:id", controllers.FindBookById)
	users.PUT("/:id", controllers.UpdateBook)
	users.PATCH("/:id", controllers.DeleteBook)
}
