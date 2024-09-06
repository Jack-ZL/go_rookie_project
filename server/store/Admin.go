package store

import (
	"time"
)

// 管理员表
type Admin struct {
	Id        int64     `grorm:"type:int(10);primary key;auto_increment" json:"id"`
	CreatedAt time.Time `grorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `grorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp" json:"updated_at"`
	Birthday  time.Time `grorm:"type:timestamp;not null;default:'1994-04-17'" json:"birthday"`
	Name      string    `grorm:"type:varchar(50);not null;default:'';comment '管理员姓名'" json:"name"`
	AdminName string    `grorm:"type:varchar(50);not null;comment '登录用户名'" json:"admin_name"`
	Password  string    `grorm:"type:varchar(40);not null;comment '密码'" json:"password"`
	Salt      string    `grorm:"type:varchar(10);not null;comment '加盐'" json:"salt"`
	Email     string    `grorm:"type:varchar(50);default:'';comment '邮箱-可用于登录'" json:"email"`
	Mobile    string    `grorm:"type:char(11);not null;default:'';comment '手机号'" json:"mobile"`
}
