package model

import "time"

// 相同类型的字段放在一起可以使内存对齐，减少内存开支
type Post struct { //binding标签代表Gin框架对请求参数的校验
	ID          int64     `json:"id,string" db:"post_id"` // 在被序列化为json时，转化成字符串格式
	AuthorID    int64     `json:"author_id,string" db:"author_id"`
	CommunityID int64     `json:"community_id,string" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"` // 帖子作者
	*Post            `json:"post"`      // 嵌入帖子详情结构体
	*CommunityDetail `json:"community"` // 嵌入社区帖子分类结构体，在结构体上加json tag会在返回时包裹整个数据返回
}
