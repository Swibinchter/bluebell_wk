package middlewares

// 用户认证的中间件

import (
	"errors"
	"fmt"
	"goWebCli/controller"
	"goWebCli/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的用户认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端一般将Token放在请求头Header(最常见)、请求体、URL中
		// 此处按照Token位于Header的Authorization中，并使用Bearer开头，具体要看业务请求
		// Authorization: Bearer xxxxxx.xx.xx / X-TOKEN: xxx.xx.xx
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Printf("获取到的请求Header是%v\n", authHeader)
		// 获取不到，返回需要登录的提示
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			// 立即终止请求处理流程，不再执行后续的代码、其他中间件和处理函数
			c.Abort()
			return
		}
		// 分割字符串，最多分割成2个字符串，获取token
		parts := strings.SplitN(authHeader, " ", 2)
		// 如果分割出来的字符串数量小于2或者开头不是Bearer
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeNeedLogin)
			// 立即终止请求处理流程，不再执行后续的代码、其他中间件和处理函数
			c.Abort()
			return
		}
		// 解析获取的token即parts[1]
		mc, err := jwt.ParseToken(parts[1])
		// 如果解析出现错误
		if err != nil {
			if errors.Is(err, errors.New("invalid token")) {
				controller.ResponseError(c, controller.CodeInvalidToken)
			} else {
				controller.ResponseError(c, controller.CodeServerBusy)
			}
			// 立即终止请求处理流程，不再执行后续的代码、其他中间件和处理函数
			c.Abort()
			return
		}
		// 验证token通过则将当前请求的userID信息保存到上下文中，后续可以通过c.GET(xx)来获取
		c.Set(controller.CtxUserIDKey, mc.UserID)

		// 将请求传递给下一个中间件或者Handler来处理，没有c.Next()，请求会一直停在当前中间件不再往后进行
		c.Next()
		// 如果c.Next()后续还有内容，则会在后续的中间件和Handler执行完之后再返回此处继续执行
	}
}
