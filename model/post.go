package model

import "time"

// 相同类型的字段放在一起可以使内存对齐，减少内存开支
type Post struct { //binding标签代表Gin框架对请求参数的校验
	ID          int64     `json:"id,string" db:"post_id"` // 在被序列化为json时，转化成字符串格式
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status" binding:""`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
