package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var Db *sqlx.DB

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
	r, err := Db.Exec("insert into person(username, sex, email)values (?, ?, ?)", "heylink", "man", "stu001@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("insert failed, ", err)
		return
	}
	fmt.Println("insert success:", id)
}

