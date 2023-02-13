package main

import (
	"ginVueBlog/model"
	"ginVueBlog/routes"
)

func main() {
	model.InitDb()      //初始化数据库
	routes.InitRouter() //初始化路由
}
