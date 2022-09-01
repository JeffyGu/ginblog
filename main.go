package main

import (
	"fmt"
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	fmt.Println("hello ")

	model.InitDb()

	routes.InitRouter()
}
