package main

import (
	"net/http"
	"io"
	"log"
)

const form = `<html><body><form action="#" method="post" name="bar">
				<input type="text" name="in"/>
				<input type="text" name="in"/>
				<input type="submit" name="Submit"/>
			</form></body></html>`

func main() {
	//浏览器输入 http://localhost:8088/test1 或http://localhost:8088/test2
	http.HandleFunc("/test1", logPanic(SimpleServer))
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
	}
}

func logPanic(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			//封装
			if x := recover(); x != nil{
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)
			}
		}()
		//真正的处理函数
		handle(w, r)
	}
}

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
	//2018/05/24 19:35:33 [[::1]:25127] caught panic: test test
	panic("test test")
}

func FormServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		//step1：解析表单数据
		r.ParseForm()
		//Form获取所有值
		io.WriteString(w, r.Form["in"][0])
		io.WriteString(w, "\n")
		io.WriteString(w, r.Form["in"][1])
		io.WriteString(w, "\n")
		//获取某个值
		io.WriteString(w, r.FormValue("in"))
	}
}
