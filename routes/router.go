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
		// 文章模块的路由接口
		router.POST("article/add", v1.AddArticle)
		router.GET("article", v1.GetArt)
		router.PUT("article/:id", v1.EditArt)
		router.DELETE("article/:id", v1.DeleteArt)

		router.GET("article/info/:id", v1.GetArtInfo) //分类下所有文章查询
		router.GET("article/list/:id", v1.GetCateArt) //单个文章查询
	}
	//启动路由,utils.HttpPort 开放端口
	r.Run(utils.HttpPort)
}
