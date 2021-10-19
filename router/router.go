package router

import (
	"fmt"
	"github.com/Peterliang233/go-blog/configs"
	"github.com/Peterliang233/go-blog/middleware"
	articleRouter "github.com/Peterliang233/go-blog/router/v1/article"
	categoryRouter "github.com/Peterliang233/go-blog/router/v1/article/category"
	commentRouter "github.com/Peterliang233/go-blog/router/v1/article/comment"
	"github.com/Peterliang233/go-blog/router/v1/article/file"
	likeRouter "github.com/Peterliang233/go-blog/router/v1/article/like"
	tagRouter "github.com/Peterliang233/go-blog/router/v1/article/tag"
	userRouter "github.com/Peterliang233/go-blog/router/v1/user"
	likeService "github.com/Peterliang233/go-blog/service/v1/api/like"
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

	// 开启一格协程定时处理redis里面的缓存的数据
	go likeService.PersistEmail()

	auth := router.Group("api/v1")
	auth.Use(middleware.JWTAuthMiddleware()) // jwt中间件认证身份信息
	{
		// 上传文件单个接口
		auth.POST("/upload", file.Upload)
		// 用户类接口
		user := auth.Group("/user")
		{
			user.PUT("/setting", userRouter.EditUser)
		}
		// 分类模块的接口
		article := auth.Group("/category")
		{
			article.POST("/add", categoryRouter.AddCategory)
			article.PUT("/:id", categoryRouter.EditCategory)
			article.DELETE("/:id", categoryRouter.DelCategory)
		}
		// 文章模块的接口
		category := auth.Group("/article")
		{
			category.POST("/add", articleRouter.AddArticle)
			category.PUT("/:id", articleRouter.EditArticle)
			category.DELETE("/:id", articleRouter.DelArticle)
		}

		// 评论模块的接口
		comment := auth.Group("/comment")
		{
			comment.POST("/add", commentRouter.AddComment)
			comment.GET("/all/:id", commentRouter.GetComment)
			comment.DELETE("/:id", commentRouter.DelComment)
		}

		// 标签模块的接口
		tag := auth.Group("/tag")
		{
			tag.POST("/add", tagRouter.AddTag)
			tag.DELETE("/:id", tagRouter.DelTag)
		}

		// 点赞模块的接口
		like := auth.Group("/like")
		{
			like.POST("", likeRouter.Like)
		}
	}
	// 获取信息的部分，这部分可以作为公共接口暴露在外面
	routerV1 := router.Group("api/v1")
	{
		routerV1.GET("/category/all", categoryRouter.GetCategory)
		routerV1.GET("/article/all", articleRouter.GetArticles)
		routerV1.GET("/article/one/:id", articleRouter.GetArticle)
		routerV1.GET("/tag/all", tagRouter.GetAllTags)
		routerV1.GET("/article/category/:id", articleRouter.GetCategoryToArticle)
		routerV1.POST("/login", userRouter.AuthHandler)
	}

	err := router.Run(configs.HttpPort)

	if err != nil {
		fmt.Println(" Listening error")
	}
}
