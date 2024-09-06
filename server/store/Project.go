package store

import (
	"time"
)

// 工作经历表
type Project struct {
	Id          int64     `grorm:"type:int(10);primary key;auto_increment" json:"id"`
	CreatedAt   time.Time `grorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `grorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp" json:"updated_at"`
	ProjectName string    `grorm:"type:varchar(100);not null;default:'';comment '项目名称'" json:"project_name"`
	FromCompany int       `grorm:"type:int(10);not null;default:0;comment '所属公司'" json:"from_company"`
	StartTime   time.Time `grorm:"type:timestamp;not null;default:'0000-00-00';comment '项目开始时间'" json:"start_time"`
	EndTime     time.Time `grorm:"type:timestamp;not null;default:'0000-00-00';comment '项目结束时间'" json:"end_time"`
	ProjectDesc string    `grorm:"type:text;not null;default:'';comment '项目描述'" json:"project_desc"`
	WorkContent string    `grorm:"type:text;not null;default:'';comment '项目负责的内容'" json:"work_content"`
}
