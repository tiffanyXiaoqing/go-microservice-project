package main

import "fmt"

func main() {
	numChan := make(chan int, 5)

	go func() {
		for i := 0; i < 50; i++ {
			numChan <- i
		}
		fmt.Println("close channel")
		close(numChan)
	}()

	for v := range numChan {
		fmt.Println("读取数据：", v)
	}

}
