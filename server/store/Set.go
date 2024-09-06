package store

import (
	"time"
)

// 设置表
type Set struct {
	Id        int64     `grorm:"type:int(10);primary key;auto_increment" json:"id"`
	CreatedAt time.Time `grorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `grorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp" json:"updated_at"`
	Key       string    `grorm:"type:varchar(64);not null;default:'';comment '设置的key'" json:"key"`
	Value     string    `grorm:"type:varchar(100);not null;default:'';comment '设置的value值'" json:"value"`
	State     int       `grorm:"type:tinyint(1);not null;default:0;comment '该设置的状态值，0-关，1-开'" json:"state"`
	Remark    string    `grorm:"type:varchar(100);not null;default:'';comment '设置的备注说明'" json:"remark"`
}
