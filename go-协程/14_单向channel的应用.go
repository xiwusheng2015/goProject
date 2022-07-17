package main

import (
	//"fmt"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// 生产
	go producter(ch)

	consumer(ch)

}

func consumer(read <-chan int) {
	for num := range read {
		time.Sleep(time.Second)
		fmt.Println("读取num = ", num)
	}
}

func producter(write chan<- int) {
	for i := 0; i < 10; i++ {
		write <- i
	}
	close(write)
}
