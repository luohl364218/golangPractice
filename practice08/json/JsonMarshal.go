package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string
	Nickname string
	Age int
}

func main()  {
	//testStruct()
	//testMap()
	//testSlice()
	testUnmarshal()
}

func testStruct() (data []byte, err error) {
	user := &User{
		Username:"heylink",
		Nickname:"hey",
		Age:18,
	}

	data, err = json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", string(data))

	return
}

func testMap() (data []byte, err error){
	var mp map[string]interface{}
	mp = make(map[string]interface{})
	mp["username"] = "heylink"
	mp["age"] = 18
	mp["sex"] = "man"

	data, err = json.Marshal(mp)
	if err != nil {
		fmt.Println(err)
		err = fmt.Errorf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
    return
}

func testSlice() {
	var se []map[string]interface{}
	mp := make(map[string]interface{})
	mp["username"] = "heylink"
	mp["age"] = 18
	mp["sex"] = "male"

	se = append(se, mp)

	mp = make(map[string]interface{})
	mp["username"] = "hexing"
	mp["age"] = 16
	mp["sex"] = "female"

	se = append(se, mp)

	data, err := json.Marshal(se)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

//反序列化
func testUnmarshal()  {

	data, err := testStruct()
	if err != nil {
		fmt.Println("get json array failed: err ", err)
		return
	}

	var user User
	//传值改变不了，一定要传指针
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json unmarshal failed: err ", err)
		return
	}
	fmt.Println(user)


	data, err = testMap()
	if err != nil {
		fmt.Println("get json array failed: err ", err)
		return
	}

	var mp map[string]interface{}
	//map在unmarshal内部分配空间，不需要手动分配
	//mp = make(map[string]interface{})
	//传指针
	err = json.Unmarshal(data, &mp)
	if err != nil {
		fmt.Println("json unmarshal failed: err ", err)
		return
	}
	fmt.Println(mp)
}