package routes

import (
	"abrigos/source/controllers"
	"abrigos/source/middleware"

	"github.com/gin-gonic/gin"
)

func bindActuatorsRoutes(router *gin.Engine) {
	router.GET("/health", controllers.Health)
	router.POST("/login", middleware.JwtMiddleware().LoginHandler)
}
