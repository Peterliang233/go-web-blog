package router

import (
	"fmt"
	"github.com/Peterliang233/go-blog/configs"
	"github.com/Peterliang233/go-blog/middleware"
	v1 "github.com/Peterliang233/go-blog/router/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() {
	gin.SetMode(configs.AppMode)
	router := gin.New() // 自定义中间件
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())   // 跨域中间件
	router.Use(middleware.Logger()) // 定义日志的中间件
	router.MaxMultipartMemory = 8
	auth := router.Group("api/v1")
	auth.Use(middleware.JWTAuthMiddleware()) // jwt中间件认证身份信息
	{
		// 上传文件单个接口
		auth.POST("/upload", v1.Upload)
		// 用户类接口
		user := auth.Group("/user")
		{
			user.PUT("/setting", v1.EditUser)
		}
		// 分类模块的接口
		article := auth.Group("/category")
		{
			article.POST("/add", v1.AddCategory)
			article.PUT("/:id", v1.EditCategory)
			article.DELETE("/:id", v1.DelCategory)
		}
		// 文章模块的接口
		category := auth.Group("/article")
		{
			category.POST("/add", v1.AddArticle)
			category.PUT("/:id", v1.EditArticle)
			category.DELETE("/:id", v1.DelArticle)
		}

		// 评论模块的接口
		comment := auth.Group("/comment")
		{
			comment.POST("/add", v1.AddComment)
			comment.GET("/all/:id", v1.GetComment)
			comment.DELETE("/:id", v1.DelComment)
		}

		// 标签模块的接口
		tag := auth.Group("/tag")
		{
			tag.POST("/add", v1.AddTag)
			tag.DELETE("/:id", v1.DelTag)
		}
	}
	// 获取信息的部分，这部分可以作为公共接口暴露在外面
	routerV1 := router.Group("api/v1")
	{
		routerV1.GET("/category/all", v1.GetCategory)
		routerV1.GET("/article/all", v1.GetArticles)
		routerV1.GET("/article/one/:id", v1.GetArticle)
		routerV1.GET("/tag/all", v1.GetAllTags)
		routerV1.GET("/article/category/:id", v1.GetCategoryToArticle)
		routerV1.POST("/login", v1.AuthHandler)
	}

	err := router.Run(configs.HttpPort)

	if err != nil {
		fmt.Println(" Listening error")
	}
}
