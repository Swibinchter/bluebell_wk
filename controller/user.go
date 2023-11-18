package controller

import (
	"errors"
	"goWebCli/dao/mysql"
	"goWebCli/logic"
	"goWebCli/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 控制器，用于参数获取、参数校验、调用业务处理、返回响应等

// SignUpHandler 用户注册
func SignUpHandler(c *gin.Context) {
	// 获取参数和校验参数(参数合法性需要在结构体中通过binding标签进行约束)
	// 创建一个结构体指针
	p := new(model.ParamSignUp)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 如果断言是validator错误则翻译中文，不是则直接返回
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}

	// 业务处理
	err = logic.SignUp(p)
	if err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 用户登录
func LoginHandler(c *gin.Context) {
	// 参数获取、参数校验
	// 获取参数和校验参数(参数合法性需要在结构体中通过binding标签进行约束)
	// 创建一个结构体指针
	p := new(model.ParamLogin)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 如果断言是validator错误则翻译中文，不是则直接返回
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}

	// 业务处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) || errors.Is(err, mysql.ErrorInvalidPassword) {
			// 用户名或密码错误
			ResponseError(c, CodeInvalidPassword)
			return
		}
		// 其他错误
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, token)

}
