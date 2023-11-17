package main

// GoWeb开发较为通用的脚手架模板

import (
	"context"
	"fmt"
	"goWebCli/controller"
	"goWebCli/dao/mysql"
	"goWebCli/dao/redis"
	"goWebCli/logger"
	"goWebCli/pkg/snowflake"
	"goWebCli/router"
	"goWebCli/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 1.加载配置，使用viper库
	if err := setting.Init(); err != nil {
		fmt.Printf("settings init failed, err:%v\n", err)
		return
	}

	// 2.初始化日志，使用zap库，接收Gin框架中的日志记录
	if err := logger.Init(setting.Config.LogConfig); err != nil {
		fmt.Printf("logger init failed, err:%v\n", err)
		return
	}
	// 将缓存中的日志写到输出位置
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	// 3.初始化连接mysql数据库
	if err := mysql.Init(setting.Config.MysqlConfig); err != nil {
		fmt.Printf("mysql init failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	zap.L().Debug("mysql init success...")

	// 4.初始化连接redis数据库
	if err := redis.Init(setting.Config.RedisConfig); err != nil {
		fmt.Printf("redis init failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	zap.L().Debug("redis init success...")

	// 初始化分布式ID生成包
	if err := snowflake.Init(setting.Config.StartTime, setting.Config.MachineID); err != nil {
		fmt.Printf("redis init failed, err:%v\n", err)
		return
	}

	// 初始化gin框架内置参数校验器的错误提示的中文翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("validator trans init failed, err:%v\n", err)
		return
	}

	// 注册路由
	r := router.SetUp()

	// 启动服务，设置优雅关机(处理完所有请求后再关机而非直接强制关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Config.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	// 创建一个接收信号的通道
	quit := make(chan os.Signal, 1)

	// signal.Notify()会转发收到的信号，第一个参数是接收的通道，后面的参数是要转发的信号类型列表
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 会发送 syscall.SIGINT 信号，常用的Ctrl+C就是这个信号
	// kill -9 会发送 syscall.SIGKILL 信号，但是不能被捕获，无需设置
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞

	// 尝试从quit中获取信号，这里会一直阻塞直到获取成功
	<-quit

	// quit中获取到信号之后就开始优雅的关机
	// 记录日志
	zap.L().Info("Shutdown server ...")
	// 创建一个5秒超时的上下文context，即5秒内如果处理不完遗留的请求也强制关机
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内关闭服务，处理遗留的请求
	err := srv.Shutdown(ctx)
	if err != nil {
		zap.L().Fatal("Server shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
