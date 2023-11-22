package mysql

import "errors"

// 定义错误类型
var (
	ErrorUserExist       = errors.New("用户名已存在")
	ErrorUserNotExist    = errors.New("用户名不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorInvalidID    = errors.New("无效的ID")
)