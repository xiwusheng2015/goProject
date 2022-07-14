package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func main() {
	go Person1()
	go Person2()
	for {

	}

}

func Person1() {
	Printer("hello")
	ch <- 666 // 给通道写入数据
}

func Person2() {
	<-ch // 读取
	Printer("world")
}

// 打印单个字符的打印机
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}
