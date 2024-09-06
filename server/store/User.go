package store

import (
	"time"
)

// user-用户表
type User struct {
	Id              int64     `grorm:"type:int(10);primary key;auto_increment" json:"id"`
	CreatedAt       time.Time `grorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt       time.Time `grorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp" json:"updated_at"`
	NickName        string    `grorm:"type:varchar(100);not null;default:'';comment '微信昵称'" json:"nick_name"`
	Avatar          string    `grorm:"type:varchar(100);not null;default:'';comment '微信头像'" json:"avatar"`
	Mobile          string    `grorm:"type:char(11);not null;default:'';index:idx_mobile;comment '手机号'" json:"mobile"`
	Openid          string    `grorm:"type:varchar(64);not null;default:''" json:"openid"`
	Unionid         string    `grorm:"type:varchar(64);not null;default:''" json:"unionid"`
	RecentLoginTime time.Time `grorm:"type:timestamp;not null;default:'0000-00-00';comment '最近登录时间'" json:"recent_login_time"`
}
