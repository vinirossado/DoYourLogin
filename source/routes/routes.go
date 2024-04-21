package routes

import (
	"doYourLogin/source/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200) // No Content
			return
		}

		c.Next()
	}
}
