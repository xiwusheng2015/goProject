package main

import "fmt"

func main() {
	ch := make(chan int) //无缓存

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //写数据
		}
		close(ch)
		ch <- 99
	}()

	for num := range ch {
		fmt.Println("num = ", num)
	}
}
