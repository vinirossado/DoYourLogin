package routes

import (
	"doYourLogin/source/controllers"
	"doYourLogin/source/middlewares"

	"github.com/gin-gonic/gin"
)

func bindActuatorsRoutes(router *gin.Engine) {
	router.GET("/health", controllers.Health)
	router.POST("/login", middlewares.JwtMiddleware().LoginHandler)
	router.POST("/webhook", controllers.HandlePost)
}
