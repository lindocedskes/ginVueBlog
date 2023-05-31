package routes

import (
	v1 "ginVueBlog/api/v1"
	"ginVueBlog/middleware"
	"ginVueBlog/utils"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "static/admin/index.html")
	p.AddFromFiles("front", "static/front/index.html")
	return p
}
func InitRouter() {
	gin.SetMode(utils.AppMode)
	//r := gin.Default() //使用默认中间件
	r := gin.New()             //new不使用任何中间件，use启用
	r.Use(middleware.Logger()) //日志中间件
	r.Use(gin.Recovery())      //该中间件从任何恐慌中恢复，并写入500（如果有）。当你的程序出现一些你未考虑到的异常时，程序就会退出，服务就停止了
	r.Use(middleware.Cors())   //使用跨域中间件

	r.HTMLRender = createMyRender()
	r.Static("/static", "./static/front/static") //前端展示，路由路径和静态资源路径
	r.Static("/admin", "./static/admin")         //后端，路由路径和静态资源路径
	r.StaticFile("/favicon.ico", "static/front/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	auth := r.Group("api/v1")       //加前缀
	auth.Use(middleware.JwtToken()) //访问路径需要使用jwt中间件的token控制。管理员才能做
	{
		//用户模块路由接口
		//auth.POST("user/add", v1.AddUser)
		//router.GET("user/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUsers)
		auth.DELETE("user/:id", v1.DeleteUsers)
		//分类模块路由接口
		auth.POST("category/add", v1.AddCate)
		//router.GET("Category", v1.GetCate)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		//router.GET("article", v1.GetArt)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

		//router.GET("article/info/:id", v1.GetArtInfo) //分类下所有文章查询
		//router.GET("article/list/:id", v1.GetCateArt) //单个文章查询

		//文件上传
		auth.POST("upload", v1.Upload)

		// 更新个人设置
		auth.PUT("profile/:id", v1.UpdateProfile)
	}

	router := r.Group("api/v1")
	{ //不需要jwt的token控制
		router.POST("user/add", v1.AddUser)
		router.GET("user/users", v1.GetUsers)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("category", v1.GetCate)
		router.GET("category/:id", v1.GetCateInfo)
		router.GET("article", v1.GetArt)
		router.GET("article/info/:id", v1.GetArtInfo) //分类下所有文章查询
		router.GET("article/list/:id", v1.GetCateArt) //单个文章查询
		router.POST("login", v1.Login)                //登陆接口

		router.GET("profile/:id", v1.GetProfile)
	}
	//启动路由,utils.HttpPort 开放端口
	r.Run(utils.HttpPort)
}
