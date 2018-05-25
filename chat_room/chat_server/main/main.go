package main

import "time"

func main() {
	initRedis("localhost:6379", 16, 1024, 300 * time.Second)

	initServer("0.0.0.0:10000")
}
