package controller

import (
	"goWebCli/logic"
	"goWebCli/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 控制器，用于参数获取、参数校验、调用业务处理、返回响应等

func SignUpHandler(c *gin.Context) {
	// 获取参数和校验参数(参数合法性需要在结构体中通过binding标签进行约束)
	// 创建一个结构体指针
	p := new(model.ParamSignUp)
	// 只能判断是否是json格式，字段名称是否正确
	err := c.ShouldBindJSON(&p)
	if err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 如果断言是validator错误则翻译中文，不是则直接返回
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			c.JSON(http.StatusOK, gin.H{"msg": errs.Translate(trans)})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		}
		return
	}

	// 业务处理
	logic.SignUp(p)

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
