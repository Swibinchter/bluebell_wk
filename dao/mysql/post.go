package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"goWebCli/model"

	"go.uber.org/zap"
)

// 新增post数据
func CreatePost(p *model.Post) (err error) {
	// 定义sql语句，用反引号包裹
	sqlStr := `insert into post(post_id, author_id, community_id, title, content) values(?,?,?,?,?)`
	// 执行sql语句, Exec执行成功会返回插入数据的主键ID
	_, err = db.Exec(sqlStr, p.ID, p.AuthorID, p.CommunityID, p.Title, p.Content)
	return err
}

// GetPostById 根据id查询post数据
func GetPostById(pid int64) (post *model.Post, err error) {
	// 定义sql语句，用反引号包裹
	sqlStr := `
			select post_id, author_id, community_id, status, title, content, create_time
			from post
			where id = ?
			`
	// 执行sql语句，返回数据保存到data中
	post = new(model.Post)
	err = db.Get(post, sqlStr, pid)
	if errors.Is(err, sql.ErrNoRows) {
		zap.L().Warn("can not find this post_id")
		err = ErrorInvalidID
	}
	fmt.Printf("查询到的post数据是%v\n", post)
	return
}
