package main

import (
	"abrigos/source/routes"
)

func main() {
	//repository.InitDB()

	route := routes.InitRouter()
	route.Run()
}
