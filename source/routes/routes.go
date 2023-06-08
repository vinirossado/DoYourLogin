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

	err := router.Run(":8025")
	if err != nil {
		panic(err)
	}

	return router
}

func bindMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	//router.Use(func(c *gin.Context) {
	//	utils.MemoryCache(c)
	//	c.Next()
	//})

	router.Use(gin.CustomRecovery(middlewares.ExceptionMiddleware))
}
