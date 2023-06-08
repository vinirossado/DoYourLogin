package main

import (
	"doYourLogin/source/repositories"
	"doYourLogin/source/routes"
	"doYourLogin/source/utils"
)

func main() {
	repositories.InitDB()

	utils.MemoryCache()

	route := routes.InitRouter()

	err := route.Run()
	if err != nil {
		return
	}

}
