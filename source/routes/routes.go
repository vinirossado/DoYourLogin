package routes

import (
	"doYourLogin/source/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// bindSwagger(router)
	bindCompanyRoutes(router)
	bindMiddlewares(router)
	bindActuatorsRoutes(router)
	bindUserRoutes(router)
	bindBookRoutes(router)
	bindAuthorRoutes(router)

	err := router.Run(":9025")
	if err != nil {
		panic(err)
	}

	return router
}

func bindMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(middlewares.ExceptionMiddleware))
}
