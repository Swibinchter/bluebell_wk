package logic

import (
	"goWebCli/dao/mysql"
	"goWebCli/model"
)

// 存放跟社区community相关的业务处理

func GetCommunityList() ([]*model.Community, error) {
	// 调用mysql的查询方法返回community的id和name
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*model.CommunityDetail, error) {
	// 调用mysql的方法返回community的详情数据
	return mysql.GetCommunityDetailByID(id)
}
