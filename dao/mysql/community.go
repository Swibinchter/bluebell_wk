package mysql

import (
	"database/sql"
	"errors"
	"goWebCli/model"

	"go.uber.org/zap"
)

// 存放于Community有关的数据库操作

// GetCommunityList 查找返回社区分类标签
func GetCommunityList() (data []*model.Community, err error) {
	// 判断表是不是为空
	sqlStr := `select count(*) from community`
	var count int
	err = db.Get(&count, sqlStr)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		zap.L().Warn("there is no community in db")
		err = nil
	}

	// 定义查询的sql语句，注意用反引号包裹
	sqlStr = `select community_id,community_name from community`
	// 结构体实例组成的切片来接收结果
	err = db.Select(&data, sqlStr)

	return
}

// GetCommunityDetailByID 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (data *model.CommunityDetail, err error) {
	// 定义sql语句，用反引号包裹
	sqlStr := `select community_id, community_name, introduction, create_time 
			from community where community_id = ?`
	// 在数据库中查询
	data = new(model.CommunityDetail)
	err = db.Get(data, sqlStr, id)
	if errors.Is(err, sql.ErrNoRows) {
		zap.L().Warn("can not find this community_id")
		err = ErrorUserIvalidID
	}
	return
}
