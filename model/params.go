package model

// 定义请求的参数结构体

// binding标签是针对validator(按照规则进行参数校验)
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
