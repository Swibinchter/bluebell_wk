package router

import (
	"goWebCli/controller"
	"goWebCli/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 路由的初始化和分配

func SetUp() *gin.Engine {
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

	return r
}
