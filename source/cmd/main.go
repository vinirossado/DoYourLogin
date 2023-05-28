package main

import (
	"doYourLogin/source/infra"
	"doYourLogin/source/repositories"
	"doYourLogin/source/routes"
)

func main() {
	repositories.InitDB()

	route := routes.InitRouter()

	//_ = utils.InitEmailServer()
	infra.InitMongo()

	err := route.Run()
	if err != nil {
		return
	}

}
