package mysql

import "goWebCli/model"

func CreatePost(p *model.Post) (err error) {
	// 定义sql语句，用反引号包裹
	sqlStr := `insert into post(post_id, author_id, community_id, title, content) values(?,?,?,?,?)`
	// 执行sql语句, Exec执行成功会返回插入数据的主键ID
	_, err = db.Exec(sqlStr, p.ID, p.AuthorID, p.CommunityID, p.Title, p.Content)
	return err
}
