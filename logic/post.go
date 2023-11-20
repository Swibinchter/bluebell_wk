package logic

import (
	"goWebCli/dao/mysql"
	"goWebCli/model"
	"goWebCli/pkg/snowflake"
)

// 存放跟帖子post相关的业务处理

// CreatePost 创建帖子
func CreatePost(p *model.Post) (err error) {
	// 判断参数中community_id是否存在
	_, err = mysql.GetCommunityDetailByID(p.CommunityID)
	if err != nil {
		return err
	}
	// 生成帖子ID
	p.ID = snowflake.GenID()
	// 调用mysql中插入对应数据的函数
	return mysql.CreatePost(p)

}
