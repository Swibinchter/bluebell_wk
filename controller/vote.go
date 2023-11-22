package controller

import (
	"goWebCli/logic"
	"goWebCli/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 投票相关请求的处理

func PostVoteHandler(c *gin.Context) {
	// 参数获取，参数校验
	p := new(model.ParamVote)
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
	// 获取当前请求的用户的id
	userId, err := GetCurrentUserId(c)
	if err != nil {
		zap.L().Error("GetCurrentUser failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}

	// 业务处理
	err = logic.VoteForPost(userId, p)
	if err != nil {
		zap.L().Error("logic.VoteForPost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, nil)
}
