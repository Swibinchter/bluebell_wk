package logic

import (
	"goWebCli/dao/mysql"
	"goWebCli/dao/redis"
	"goWebCli/model"
	"goWebCli/pkg/snowflake"

	"go.uber.org/zap"
)

// 存放跟帖子post相关的业务处理

// CreatePost 创建帖子
func CreatePost(p *model.Post) (err error) {
	// 判断参数中community_id是否存在
	_, err = mysql.GetCommunityDetailById(p.CommunityID)
	if err != nil {
		return err
	}
	// 生成帖子ID
	p.ID = snowflake.GenID()
	// 调用mysql中插入对应数据的函数
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	// redis创建帖子
	err = redis.CreatePost(p.ID, p.CommunityID)
	return

}

// GetPostById 根据id获取帖子详情
func GetPostById(pid int64) (data *model.ApiPostDetail, err error) {
	data = new(model.ApiPostDetail)
	// 调用mysql方法查询id对应的记录
	data.Post, err = mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	// 根据返回的author_id查询用户信息
	user, err := mysql.GetUserById(data.Post.AuthorID)
	data.AuthorName = user.UserName
	if err != nil {
		zap.L().Error("mysql.GetUserById failed", zap.Int64("author_id", data.Post.AuthorID), zap.Error(err))
		return
	}

	// 根据返回的community_id查询帖子分类信息
	data.CommunityDetail, err = mysql.GetCommunityDetailById(data.Post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("community_id", data.Post.CommunityID), zap.Error(err))
		return
	}
	return
}
