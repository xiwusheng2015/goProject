package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)

	fmt.Printf("len(ch) = %d,cap(ch) = %d\n", len(ch), cap(ch))

	// 新建一个协程
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子协程： i= ", i)
			ch <- i // 往通道写内容，如果写入的数据没有被读取，便会阻塞

		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("num = ", num)
	}
}
