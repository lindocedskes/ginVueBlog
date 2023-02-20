package routes

import (
	v1 "ginVueBlog/api/v1"
	"ginVueBlog/utils"
)
import "github.com/gin-gonic/gin"

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1") //加前缀
	{
		//用户模块路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("user/users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUsers)
		router.DELETE("user/:id", v1.DeleteUsers)
		//分类模块路由接口
		router.POST("Category/add", v1.AddCate)
		router.GET("Category", v1.GetCate)
		router.PUT("Category/:id", v1.EditCate)
		router.DELETE("Category/:id", v1.DeleteCate)
	}
	//启动路由,utils.HttpPort 开放端口
	r.Run(utils.HttpPort)
}
