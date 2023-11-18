package logic

import (
	"errors"
	"goWebCli/dao/mysql"
	"goWebCli/model"
	"goWebCli/pkg/jwt"
	"goWebCli/pkg/snowflake"
)

// 根据业务逻辑执行调用不同的工具包或者数据库操作函数

// SignUp 用户注册的业务处理
func SignUp(p *model.ParamSignUp) (err error) {
	// 判断用户是否存在
	_, err = mysql.CheckUserExist(p.Username)
	if err == nil {
		return mysql.ErrorUserExist
	} else if !errors.Is(err, mysql.ErrorUserNotExist) {
		return err
	}

	// 生成UID，使用雪花算法生成分布式ID
	userID := snowflake.GenID()

	// 构造一个结构体实例保存user信息
	user := &model.User{
		UserID:   userID,
		UserName: p.Username,
		Password: p.Password,
	}
	// 保存到数据库
	return mysql.InsertUser(user)
}

// Login 用户登录的业务处理
func Login(p *model.ParamLogin) (token string, err error) {
	// 查询用户名是否存在
	user, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return
	}

	// 对比用户密码是否正确
	err = mysql.ValidatePassWord(p.Password, user.Password)
	if err != nil {
		return
	}

	// 生成并返回JWT
	return jwt.GenToken(user.UserID, user.UserName)
}
