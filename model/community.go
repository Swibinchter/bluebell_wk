package model

import "time"

type Community struct {
	// go中int64类型的数值范围，大于 前端js语言的数字类型表示的范围，可能会出现超范围导致数字失真情况
	// 故而需要将int64类型以字符串形式传递给前端，只需要在json tag中加上string即可
	ID   int64  `json:"id,string" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id,string" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"` // omitempty表示允许为空，json遇到空时忽略这个字符串
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
