package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	// 消费者
	go func() {
		for i := 0; i < 15; i++ {
			num := <-ch
			fmt.Println("读取->", num)
		}
		quit <- true
	}()

	// 生产者
	fibonacci(ch, quit)
}

func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}
