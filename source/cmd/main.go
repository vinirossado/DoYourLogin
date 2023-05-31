package main

import (
	"doYourLogin/source/repositories"
	"doYourLogin/source/routes"
)

func main() {
	repositories.InitDB()

	route := routes.InitRouter()

	//_ = utils.InitEmailServer()
	//go build -gcflags -m=2
	err := route.Run()
	if err != nil {
		return
	}

}
