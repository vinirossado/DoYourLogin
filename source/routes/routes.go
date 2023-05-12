package routes

import (
	"abrigos/source/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// bindSwagger(router)
	bindMiddlewares(router)
	bindActuatorsRoutes(router)
	bindUserRoutes(router)

	err := router.Run(":8025")
	if err != nil {
		panic(err)
	}

	return router
}

func bindMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(middleware.ExceptionMiddleware))
}
