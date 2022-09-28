package main

import (
	"ginblog/model"
	"ginblog/routes"
)

// @title ginblog系统
// @version 1.0
// @description ginblog项目
// @termsOfService https://github.com/jeffygu/ginblog
func main() {

	model.InitDb()

	routes.InitRouter()
}
