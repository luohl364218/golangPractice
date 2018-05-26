package main

import "fmt"

//客户端交互的逻辑都在这里
func logic()  {
	enterMenu()
}

func enterMenu()  {
	fmt.Println("1.list online user")
	fmt.Println("2.talk")
	fmt.Println("3.exit")

	var sel int
	fmt.Scanf("%d", &sel)
	switch sel {
	case 1:
		showOnlineUserList()
	case 2:
		enterTalk()
	case 3:
		return
	}
}

func enterTalk() {

}