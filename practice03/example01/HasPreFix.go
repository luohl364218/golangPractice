package main

import (
	"fmt"
	"strings"
)
/*strings的使用*/
func main()  {
	var (
		url string
		path string
	)

	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)
	path = pathProcess(path)

	fmt.Println(url)
	fmt.Println(path)
}

//判断URL是否以http开头，不是则加上
func urlProcess(url string)  string{
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}
//判断路径是否以/结尾，不是则加上
func pathProcess(path string) string {
	result := strings.HasSuffix(path,"/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}



