package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name string  `json:"name_param"`
	Age int  `json:"age"`
	Score float32 `json:"score"`
}


func main()  {
	//初始化
	var stu Stu
	stu.Name = "heylink"
	stu.Age = 25
	stu.Score = 99.9
	//打包 通过反射拿到tag
	data, err := json.Marshal(stu)
	if err != nil {
		return
	}

	fmt.Printf(string(data))

}


