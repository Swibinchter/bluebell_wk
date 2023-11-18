package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
封装返回响应的函数，格式如下
{
	"code": 1000, 	// 程序中的错误代码
	"msg": xx, 		// 提示信息
	"data": {} 		// 数据信息
}
*/

// 定义错误代码常量
const (
	CodeSuccess int64 = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
)

// 定义错误提示信息
var codeMsgMap = map[int64]string{
	CodeSuccess:         "Success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效的Token",
}

type ResponseData struct {
	Code int64       `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// 根据错误代码查询提示信息
func GetMsg(code int64) string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

// 返回成功响应
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  GetMsg(CodeSuccess),
		Data: data,
	})
}

// 返回错误响应及约定的提示信息
func ResponseError(c *gin.Context, code int64) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  GetMsg(code),
		Data: nil,
	})
}

// 返回错误响应，及定制的提示信息
func ResponseErrorWithMsg(c *gin.Context, code int64, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
