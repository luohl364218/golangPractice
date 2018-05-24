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
	_, err := Db.Exec("update person set username = ? where user_id=?;", "heylinkNew", 1)
	if err != nil {
		fmt.Println("update failed, ", err)
		return
	}
	fmt.Println("update success")
}

