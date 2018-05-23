package main

import "testing"

//文件名字要为 文件名+ _test
//方法名要用Test开头
func TestAdd(t *testing.T) {
	r := Add(2, 4)
	if r != 6 {
		t.Fatalf("add(2,4) error, expect:%d, actual:%d", 6, r)
	}
	t.Log("test add succ")
}
