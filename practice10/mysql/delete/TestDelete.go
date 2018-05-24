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
	_, err := Db.Exec("delete from person where user_id=?", 1)
	if err != nil {
		fmt.Println("delete failed, ", err)
		return
	}
	fmt.Println("delete success")
}