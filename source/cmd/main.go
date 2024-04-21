package main

import (
	"doYourLogin/source/repositories"
	"doYourLogin/source/routes"
	cors "github.com/gin-contrib/cors"
)

func main() {
	repositories.InitDB()

	route := routes.InitRouter()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	route.Use(cors.New(config))

	err := route.Run()
	if err != nil {
		return
	}

}
