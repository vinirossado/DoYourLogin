package main

import (
	"doYourLogin/source/repositories"
	"doYourLogin/source/routes"
)

func main() {
	repositories.InitDB()

	route := routes.InitRouter()

	//_ = utils.InitEmailServer()

	err := route.Run()
	if err != nil {
		return
	}

}
