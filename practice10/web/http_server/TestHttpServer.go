package main

import (
	"net/http"
	"fmt"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello ")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	//读取参数
	userName := r.PostFormValue("userName")
	fmt.Println(userName)
	fmt.Fprintf(w, "welcome %s", userName)
}

func main() {
	//设置相应处理函数
	http.HandleFunc("/", Hello)
	http.HandleFunc("/login", Login)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}


