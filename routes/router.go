package routes

import (
	v1 "ginblog/api/v1"
	_ "ginblog/docs"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()

	//注册跨越
	r.Use(middleware.Cors())
	//注册日志中间件
	r.Use(middleware.Logger())

	url := ginSwagger.URL("http://127.0.0.1:3000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	auth := r.Group("api/v1")
	//鉴权路由
	auth.Use(middleware.JwtToken())
	{
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		// 上传文件
		auth.POST("upload", v1.UpLoad)
	}
	//不需要鉴权路由
	router := r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)

		// 分类模块的路由接口
		router.GET("categorys", v1.GetCategorys)

		// 文章模块的路由接口
		router.GET("articles", v1.GetArticles)
		router.GET("article/info/:id", v1.GetArtInfo)

		router.POST("login", v1.Login)

	}
	r.Run(utils.HttpPort)
}
