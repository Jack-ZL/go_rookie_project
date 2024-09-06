package store

import (
	"time"
)

// 菜单表
type Menu struct {
	Id        int64     `grorm:"type:int(10);primary key;auto_increment" json:"id"`
	CreatedAt time.Time `grorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `grorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp" json:"updated_at"`
	Level     int       `grorm:"type:tinyint(1);not null;default:1" json:"level"`
	ParentId  int64     `grorm:"type:int(10);not null;default:0" json:"parent_id"`
	Icon      string    `grorm:"type:varchar(50);not null;default:''" json:"icon"`
	Name      string    `grorm:"type:varchar(50);not null;default:''" json:"name"`
	Path      string    `grorm:"type:varchar(100);not null;default:''" json:"path"`
	Desc      string    `grorm:"type:varchar(100);not null;default:''" json:"desc"`
	Order     int       `grorm:"type:tinyint(2);not null;default:0;comment '排序：越小越靠前'" json:"order"`
	IsShow    int       `grorm:"type:tinyint(1);not null;default:0;comment '是否显示：0-不显示，1-显示'" json:"is_show"`
}
