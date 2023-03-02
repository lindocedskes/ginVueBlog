package routes

import (
	v1 "ginVueBlog/api/v1"
	"ginVueBlog/middleware"
	"ginVueBlog/utils"
)
import "github.com/gin-gonic/gin"

func InitRouter() {
	gin.SetMode(utils.AppMode)
	//r := gin.Default() //使用默认中间件
	r := gin.New()                  //new不使用任何中间件，use启用
	r.Use(middleware.Logger())      //自定义中间件
	r.Use(gin.Recovery())           //该中间件从任何恐慌中恢复，并写入500（如果有）。当你的程序出现一些你未考虑到的异常时，程序就会退出，服务就停止了
	auth := r.Group("api/v1")       //加前缀
	auth.Use(middleware.JwtToken()) //访问路径需要使用jwt中间件的token控制。管理员才能做
	{
		//用户模块路由接口
		//auth.POST("user/add", v1.AddUser)
		//router.GET("user/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUsers)
		auth.DELETE("user/:id", v1.DeleteUsers)
		//分类模块路由接口
		auth.POST("Category/add", v1.AddCate)
		//router.GET("Category", v1.GetCate)
		auth.PUT("Category/:id", v1.EditCate)
		auth.DELETE("Category/:id", v1.DeleteCate)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		//router.GET("article", v1.GetArt)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

		//router.GET("article/info/:id", v1.GetArtInfo) //分类下所有文章查询
		//router.GET("article/list/:id", v1.GetCateArt) //单个文章查询

		//文件上传
		auth.POST("upload", v1.Upload)
	}

	router := r.Group("api/v1")
	{ //不需要jwt的token控制
		router.POST("user/add", v1.AddUser)
		router.GET("user/users", v1.GetUsers)
		router.GET("Category", v1.GetCate)
		router.GET("article", v1.GetArt)
		router.GET("article/info/:id", v1.GetArtInfo) //分类下所有文章查询
		router.GET("article/list/:id", v1.GetCateArt) //单个文章查询
		router.POST("login", v1.Login)                //登陆接口
	}
	//启动路由,utils.HttpPort 开放端口
	r.Run(utils.HttpPort)
}
