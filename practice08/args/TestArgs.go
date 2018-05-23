package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("length of args:", len(os.Args))
	for i, v := range os.Args{
		fmt.Printf("args[%d] = %s\n", i, v)
	}

}

