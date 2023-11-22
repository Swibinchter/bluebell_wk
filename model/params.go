package model

// 定义请求的参数结构体
// binding标签是针对validator(按照规则进行参数校验)

// ParamSignUp 注册时请求的参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录时请求的参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVote 投票参数
type ParamVote struct {
	PostId string `json:"post_id" binding:"required"` // 帖子id
	Direction int8 `json:"direction,string" binding:"oneof=1 0 -1"` // 投票类型，oneof限制只能是 1赞成票 0取消投票 -1反对票
}
