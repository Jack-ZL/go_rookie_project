package db

import (
	"fmt"
	"github.com/Jack-ZL/go_rookie/config"
	"github.com/Jack-ZL/go_rookie/orm"
	"net/url"
	"strconv"
	"time"
)

type BaseModel struct {
	Id        int64     `grorm:"type:bigint(20),primary key,auto_increment" json:"id"`
	CreatedAt time.Time `grorm:"type:timestamp,not null,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `grorm:"type:timestamp,not null,default:current_timestamp on update current_timestamp" json:"updated_at"`
}

// 公共db模型
var DB *orm.GrDb

/*
 * 自动调用init函数，只会被调用一次
 */
func init() {
	mysqlConfig := config.Conf.Mysql
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&loc=%s&parseTime=true",
		mysqlConfig["db_user"],
		mysqlConfig["db_password"],
		mysqlConfig["db_host"],
		mysqlConfig["db_port"],
		mysqlConfig["db_database"],
		mysqlConfig["db_charset"],
		url.QueryEscape("Asia/Shanghai"))
	DB = orm.Open("mysql", dataSourceName)
	conns, _ := strconv.Atoi(mysqlConfig["db_max_conn"].(string))
	DB.SetMaxIdleConns(conns)
	DB.Prefix = mysqlConfig["db_prefix"].(string)
}
