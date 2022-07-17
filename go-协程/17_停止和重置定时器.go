package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)

	isOk := timer.Reset(1 * time.Second)

	fmt.Println(isOk)

	go func() {
		<-timer.C
		fmt.Println("子协程开始执行")
	}()

	for {

	}

}
