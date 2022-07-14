package main

import (
	"fmt"
	"time"
)

func main() {
	go test()

	for {
		fmt.Println("this is main")
		time.Sleep(time.Second * 2)
	}
}

func test() {
	for {
		fmt.Println("this is test")
		time.Sleep(time.Second)
	}
}
