package model

import (
	"fmt"

	"github.com/gxin0426/beegolearn/utils"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int
	Createtime int64
}

//----------数据库操作----------------

//insert
func InsertUser(user User) (int64, error) {

	return utils.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

//get

func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

// get id for username

func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

// get id for username and pwd

func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}
