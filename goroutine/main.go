package main

import (
	"fmt"
	"time"
)

func display() {
	count := 1
	for {
		fmt.Println("======> 这是go协程", count)
		count++
	}
}

func main() {
	//启动了go程
	go display()

	count := 1
	for {
		fmt.Println("这是主go程", count)
		count++
		time.Sleep(1 * time.Second)
	}
}
