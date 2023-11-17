package logic

import (
	"goWebCli/dao/mysql"
	"goWebCli/model"
)

// 根据业务逻辑执行调用不同的工具包或者数据库操作函数

func SignUp(p *model.ParamSignUp) {
	// 判断用户是否存在
	// 生成UID
	// 保存到数据库
	mysql.InsertUser()
}
