package main

import (
	"doYourLogin/source/repositories"
	"doYourLogin/source/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	repositories.InitDB()
	route := routes.InitRouter()
	router := gin.Default()
	router.Run()
	err := route.Run()
	if err != nil {
		return
	}

}
