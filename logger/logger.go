package logger

// 使用zap记录日志，并接收Gin框架中的日志记录

import (
	"goWebCli/setting"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Init初始化Logger
func Init(cfg *setting.LogConfig) (err error) {
	// 设置日志的编码
	encoder := getEncoder()

	// 设置日志的输出位置，分割备份信息
	writeSyncer := getLogWriter(
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
	)

	// 设置日志的过滤级别
	// 将设置的level字符串转换成zap库能识别的格式，就相当于根据配置文件将l赋值为zapcore.DebugLevel等对应级别
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return err
	}

	// 创建一个新的日志器
	core := zapcore.NewCore(encoder, writeSyncer, l)
	lg := zap.New(core, zap.AddCaller())

	// 替换zap默认的全局日志器，其他包中即可调用zap.L()
	zap.ReplaceGlobals(lg)
	return
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	})
}

func getEncoder() zapcore.Encoder {
	// 先创建默认配置
	encoderConfig := zap.NewProductionEncoderConfig()

	// 时间显示格式，从时间戳改为可读性更强的时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 时间格式显示到秒
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 将日期信息存放的键值对中的key设置为time
	encoderConfig.TimeKey = "time"
	// 日志级别用大写字母表示
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 调用函数显示的格式设置为短格式，即减少文件路径的显示
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// 使用JSON编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// Gin框架默认中间件Logger()会将日志输出到标准输出如终端命令行，使用Zap来替换它，就可以实现Gin日志的接收
// 重写两个中间件Logger()和Recovery()，后续在创建Gin的router时调用它们
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// 重写中间件Recovery()
// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
