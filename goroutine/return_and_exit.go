package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//启动了go程
	go func() {
		func() {
			fmt.Println("1")
			//return  返回当前函数
			os.Exit(-1) // 返回当前进程
			// runtime.Goexit() //推出当前子协程
		}()

		fmt.Println("2")
	}()

	fmt.Println("3")
	time.Sleep(5 * time.Second)
	fmt.Println("4")
}
