package routers

import (
	"fmt"
	v1 "github.com/Peterliang233/go-blog/api/v1"
	"github.com/Peterliang233/go-blog/middleware"
	"github.com/Peterliang233/go-blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.New() //自定义中间件
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())   //跨域中间件
	router.Use(middleware.Logger()) //定义日志的中间件
	router.MaxMultipartMemory = 8
	auth := router.Group("api/v1")
	auth.Use(middleware.JWTAuthMiddleware()) //jwt中间件认证身份信息
	{
		//上传文件单个接口
		auth.POST("/upload", v1.Upload)
		//用户模块的接口
		user := auth.Group("/user")
		{
			user.POST("/add", v1.AddUser)
			user.PUT("/:id", v1.EditUser)
			user.DELETE("/:id", v1.DelUser)
		}
		//分类模块的接口
		article := auth.Group("/category")
		{
			article.POST("/add", v1.AddCategory)
			article.PUT("/:id", v1.EditCategory)
			article.DELETE("/:id", v1.DelCategory)
		}
		//文章模块的接口
		category := auth.Group("/article")
		{
			category.POST("/add", v1.AddArticle)
			category.PUT("/:id", v1.EditArticle)
			category.DELETE("/:id", v1.DelArticle)
		}
	}
	//获取信息的部分，这部分可以作为公共接口暴露在外面
	routerV1 := router.Group("/api/v1")
	{
		routerV1.GET("/user/search", v1.GetUsers)
		routerV1.GET("/category/search", v1.GetCategory)
		routerV1.GET("/article/search", v1.GetArticles)
		routerV1.GET("/article/one/:id", v1.GetArticle)
		routerV1.GET("/article/category/:id", v1.GetCategoryToArticle)
		routerV1.POST("/login", v1.AuthHandler) //登录接口
	}
	err := router.Run(utils.HttpPort)
	if err != nil {
		fmt.Println(" Listening error")
	}
}
