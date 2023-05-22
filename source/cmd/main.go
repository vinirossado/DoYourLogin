package main

import (
	"doYourLogin/source/repositories"
	"doYourLogin/source/routes"
)

func main() {
	repositories.InitDB()

	route := routes.InitRouter()

	//var _ = utils.InitEmailServer()

	route.Run()

}
