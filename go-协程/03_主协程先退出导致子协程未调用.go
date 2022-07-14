package main

import (
	"fmt"
	"time"
)

func main() {
	go func(){
		for {
			fmt.Println("child goroutine")
			time.Sleep(time.Second)
		}
	}()
	fmt.Println("main end..")
}
