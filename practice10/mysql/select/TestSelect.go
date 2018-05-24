package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var Db *sqlx.DB

type Person struct {
	UserId int `db:"user_id"`
	UserName string `db:"username"`
	Sex string `db:"sex"`
	Email string `db:"email"`
}

func init() {
	//root是数据库用户名，admin123是密码
	database, err := sqlx.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("open mysql failed, ", err)
		return
	}
	Db = database
}

func main() {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 1)
	if err != nil {
		fmt.Println("select failed, ", err)
		return
	}

	fmt.Println("select success:", person)
}

