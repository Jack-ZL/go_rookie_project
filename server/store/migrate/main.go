package main

import (
	"goRookie_project/common/db"
	"goRookie_project/server/store"
)

/*
生成表
在当前目录命令行直接运行：go run main
*/
func main() {
	//data := &store.Admin{
	//	Id:        1,
	//	CreatedAt: time.Now(),
	//	UpdatedAt: time.Now(),
	//	Birthday:  time.Now(),
	//	Name:      "赵龙",
	//	AdminName: "Jack-Z",
	//	Password:  "123456",
	//	Salt:      "jackzl",
	//	Email:     "zlitmh@outlook.com",
	//	Mobile:    "15738869736",
	//}
	//insert, i, err := db.DB.New(&store.Admin{}).Insert(data)
	//fmt.Println(i, insert, err)

	db.DB.AutoMigrateMySQL(&store.Admin{})
	db.DB.AutoMigrateMySQL(&store.Menu{})
	db.DB.AutoMigrateMySQL(&store.Project{})
	db.DB.AutoMigrateMySQL(&store.Set{})
	db.DB.AutoMigrateMySQL(&store.User{})
	db.DB.AutoMigrateMySQL(&store.WorkExperience{})
}
