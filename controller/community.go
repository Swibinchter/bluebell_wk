package controller

import (
	"goWebCli/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用来处理社区相关的路由

// CommunityHandler 处理请求社区分类标签的路由
func CommunityHandler(c *gin.Context) {
	// 获取参数、校验参数（此处无需）

	// 业务处理
	// 查询所有社区的id和name并以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		// 一般除了跟前端约定好的普通错误，后端出现的错误都不轻易暴露，只在前端显示服务忙之类的
		// 后端自己记录错误到日志
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, data)

}

// CommunityDetailHandler 处理请求社区标签详情的路由
func CommunityDetailHandler(c *gin.Context) {
	// 获取参数，url中的参数
	idStr := c.Param("id")
	// 校验参数，将字符串转换成int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get community_id failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 业务处理
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, data)
}
