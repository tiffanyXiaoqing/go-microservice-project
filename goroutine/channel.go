package main

import (
	"fmt"
	"time"
)

func main() {
	//sync.RWMutex{}
	//A往通道协数据，B往通道读数据

	//创建逛到，装数字
	numChan := make(chan int, 5)
	//读5个，写5个
	//strChan := make(chan string) //装字符串

	//创建两个go程，父亲写数据，儿子读数据
	go func() {
		for i := 0; i < 60; i++ {
			data := <-numChan
			fmt.Println("son1 read data:", data)
		}
	}()

	for i := 1; i < 50; i++ {
		numChan <- i
		fmt.Println("=====>这是主go程", i)
	}

	time.Sleep(5 * time.Second)
}
