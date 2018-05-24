package main

import (
	"html/template"
	"fmt"
	"net/http"
	"os"
)

type Person struct {
	Name string
	Age int
}

func main() {
	//终端输出
	//TestTemplate()

	////浏览器输入 http://localhost:8088/test
	http.HandleFunc("/test", SimpleServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
	}
}

//在终端输出
func TestTemplate()  {
	temp, err := template.ParseFiles("./practice10/web/template/index.html")
	if err != nil {
		fmt.Println("parse file error",err)
		return
	}
	p := Person{
		Name:"heylink",
		Age:18,
	}

	if err = temp.Execute(os.Stdout, p); err != nil {
		fmt.Println("execute template error", err)
	}
}

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./practice10/web/template/index.html")
	if err != nil {
		fmt.Println("parse file error",err)
		return
	}
	p := Person{
		Name:"heylink",
		Age:18,
	}

	if err = temp.Execute(w, p); err != nil {
		fmt.Println("execute template error", err)
	}
}

