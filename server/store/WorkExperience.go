package store

import (
	"time"
)

// 工作经历表
type WorkExperience struct {
	Id           int64     `grorm:"type:int(10);primary key;auto_increment" json:"id"`
	CreatedAt    time.Time `grorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `grorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp" json:"updated_at"`
	Company      string    `grorm:"type:varchar(100);not null;default:'';comment '公司名'" json:"company"`
	CompanyType  string    `grorm:"type:varchar(30);not null;default:'';comment '公司类型'" json:"company_type"`
	StartTime    time.Time `grorm:"type:timestamp;not null;default:'0000-00-00';comment '入职时间'" json:"start_time"`
	EndTime      time.Time `grorm:"type:timestamp;not null;default:'0000-00-00';comment '离职时间'" json:"end_time"`
	PositionHeld string    `grorm:"type:varchar(20);not null;default:'';comment '担任职位'" json:"position_held"`
	Location     string    `grorm:"type:varchar(100);not null;default:'';comment '所在地'" json:"location"`
	ComDesc      string    `grorm:"type:text;not null;default:'';comment '公司描述'" json:"com_desc"`
	PositionDesc string    `grorm:"type:text;not null;default:'';comment '所在职位负责工作'" json:"position_desc"`
}
