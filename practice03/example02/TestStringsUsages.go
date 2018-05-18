package main

import (
	"strings"
	"fmt"
	"strconv"
)
//strings strconv基本用法
func main()  {

	str := "   hello world abc  "
	result := strings.Replace(str, "world", "you", 1)
	fmt.Println("replace:", result)

	count := strings.Count(str, "l")
	fmt.Println("count:", count)

	result = strings.Repeat(str, 3)
	fmt.Println("repeat:", result)

	result = strings.ToLower(str)
	fmt.Println("tolower:", result)

	result = strings.ToUpper(str)
	fmt.Println("toUpper", result)

	result = strings.TrimSpace(str)
	fmt.Println("trimspace", result)

	result = strings.Trim(str, " \n\r")
	fmt.Println("trim:", result)

	result = strings.TrimLeft(str, " \n\r")
	fmt.Println("trimLeft:", result)

	result = strings.TrimRight(str, " ab")
	fmt.Println("trimRight:", result)

	splitResult := strings.Fields(str)
	for i := 0; i < len(splitResult); i++ {
		fmt.Println("fields:", splitResult[i])
	}

	splitResult = strings.Split(str, "l")
	for i := 0; i < len(splitResult); i++ {
		fmt.Println("split:", splitResult[i])
	}

	str2 := strings.Join(splitResult, "l")
	fmt.Println("join:", str2)

	str2 = strconv.Itoa(1000)
	fmt.Println("itoa:", str2)

	number, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("cannot convert to int", err)
		return
	}
	fmt.Println("atoi:", number)
}

