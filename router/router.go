package router

import (
	"goWebCli/logger"
	"goWebCli/setting"
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
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, setting.Config.Version)
	})

	return r
}
