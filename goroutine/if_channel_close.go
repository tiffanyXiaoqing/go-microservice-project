package main

import "fmt"

func main() {
	numChan := make(chan int, 5)

	go func() {
		for i := 0; i < 10; i++ {
			numChan <- i
		}

		close(numChan)
	}()

	for {
		if val, ok := <-numChan; !ok {
			fmt.Println("管道已经关闭，准备退出2")
			break
		} else {
			msg := fmt.Sprintf("接收%v\n", val)
			fmt.Println(msg)
		}
	}
}
