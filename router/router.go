package router

import (
	"goWebCli/controller"
	"goWebCli/logger"
	"goWebCli/middleware"

	"github.com/gin-gonic/gin"
)

// 路由的初始化和分配

func SetUp(mode string) *gin.Engine {
	// 当配置文件里的mode是发布模式，设置gin的模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建一个新engine，不使用默认的default
	r := gin.New()
	// 模仿default启用两个中间件，以便将日志输出给zap
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 分配路由
	// 业务路由

	// 定义一个路由组
	v1 := r.Group("/api/v1")

	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)

	// 登录后才能访问的路由
	// 调用JWT认证中间件
	v1.Use(middleware.JWTAuthMiddleware())
	{
		// 查询社区的帖子分类
		v1.GET("/community",controller.CommunityHandler)
		// 根据id查询社区分类标签的详情，通过路径传递参数id
		v1.GET("community/:id",controller.CommunityDetailHandler)
		
		// 创建帖子
		v1.POST("/post", controller.CreatePostHandler)


	}

	return r
}
