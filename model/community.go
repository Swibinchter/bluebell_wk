package model

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introductiong,omitempty" db:"introduction"` // omitempty表示允许为空，json遇到空时忽略这个字符串
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
