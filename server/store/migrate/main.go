package main

import (
	"encoding/json"
	"fmt"
	"goRookie_project/common/db"
)

// 假设我们有一个User结构体，它包含了一些字段和索引信息
type UserTest struct {
	ID       int64  `grorm:"type:int,primary key,auto_increment" json:"id"`
	Name     string `grorm:"type:varchar(30),index:idx_name" json:"name"`
	Email    string `grorm:"type:varchar(30),unique_index:idx_email" json:"email"`
	Age      int    `grorm:"type:int(3)" json:"age"`
	Password string `grorm:"type:varchar(50)" json:"password"`
}

// 生成表
func main() {
	data := &UserTest{}
	anies, err := db.DB.New(&UserTest{}).Select(data)

	marshal, err := json.Marshal(anies)
	fmt.Println(string(marshal))
	fmt.Println(err)
}
