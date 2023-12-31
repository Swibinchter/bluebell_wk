package controller

import (
	"goWebCli/logic"
	"goWebCli/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 与帖子有关的控制器

// CreatePostHandler 处理创建帖子的请求
func CreatePostHandler(c *gin.Context) {
	// 获取参数，校验参数，通过gin自带校验器和字段后的binding标签
	p := new(model.Post)
	err := c.ShouldBindJSON(p)
	if err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("create post with invalid param", zap.Error(err))
		// 如果断言是validator错误则翻译中文，不是则直接返回
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}

	// 从当前上下文中获取用户id
	p.AuthorID, err = GetCurrentUserId(c)
	if err != nil {
		zap.L().Error("GetCurrentUser failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}

	// 业务处理

	err = logic.CreatePost(p)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 处理通过id获取帖子详情的请求
func GetPostDetailHandler(c *gin.Context) {
	// 获取参数、校验参数
	// 获取url中的字符串id
	pidStr := c.Param("id")
	// 转换成int类型
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("Get post id failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 业务处理
	data, err:=logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, data)
}
