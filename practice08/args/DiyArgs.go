package main

import (
	"flag"
	"fmt"
)

func main() {

	var classpath string
	var logLevel int
	//在配置项中加入参数 -c c:/test.conf -d 5
	flag.StringVar(&classpath, "c", "", "please input class path")
	flag.IntVar(&logLevel, "d", 10, "define log level")
	//生效
	flag.Parse()

	fmt.Println("path:", classpath)
	fmt.Println("log level:", logLevel)
}
