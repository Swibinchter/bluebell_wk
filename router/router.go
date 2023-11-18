package router

import (
	"goWebCli/controller"
	"goWebCli/logger"
	"goWebCli/middlewares"
	"net/http"

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
	// 当请求没有匹配到任何路径时的处理
	r.NoRoute(func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"msg": "404"}) })
	// 业务路由
	// 注册
	r.POST("/signup", controller.SignUpHandler)
	// 登录
	r.POST("/login", controller.LoginHandler)
	// 测试
	r.POST("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		controller.ResponseSuccess(c,nil)
	})

	return r
}
