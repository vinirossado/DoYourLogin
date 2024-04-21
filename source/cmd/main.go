package main

import (
	"doYourLogin/source/repositories"
	"doYourLogin/source/routes"
)

func main() {
	repositories.InitDB()

	route := routes.InitRouter()

	err := route.Run()
	if err != nil {
		return
	}

}
