package mysql

import (
	"database/sql"
	"errors"
	"goWebCli/model"
	"goWebCli/pkg/md5"
	"goWebCli/setting"
)

// 将每一步数据库的操作封装成函数，让logic/server层依据业务来调用

// CheckUserExist 检查对应用户名是否存在，不存在的话返回nil，存在的话返回user结构体实例
func CheckUserExist(username string) (user model.User, err error) {
	// 编写查询用户名对应的语句，用反引号包裹确保语句不被转义
	sqlStr := `select user_id, username, password from user where username = ?`
	// 获取结果存入user结构体实例中
	err = db.Get(&user, sqlStr, username)
	if errors.Is(err, sql.ErrNoRows) {
		// 如果查不到用户信息，则返回提示信息
		return user, ErrorUserNotExist
	} else if err != nil {
		// 其他错误直接返回
		return
	}
	// 如果能查到用户名，返回结构体实例user，err为nil
	return user, nil
}

// InsertUser 向数据库中插入一条新用户记录
func InsertUser(user *model.User) (err error) {
	// 对密码加密，使用md5算法和配置文件中的salt
	user.Password = md5.EncryptPassword(user.Password, setting.Config.Salt)

	// 保存到数据库
	// 编写sql语句，用反引号包裹确保语句不被转义
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, user.Password)
	return
}

func ValidatePassWord(plainPassword, password string) (err error) {
	// 验证密码，使用md5算法和配置文件中的salt
	if !md5.ValidatePassWord(plainPassword, setting.Config.Salt, password) {
		return ErrorInvalidPassword
	}
	return
}
