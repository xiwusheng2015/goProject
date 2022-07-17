package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(2*time.Second)
	<- timer.C
	fmt.Println("时间到")

	time.Sleep(2* time.Second)
	fmt.Println("时间到02")

	<- time.After(2*time.Second)
	fmt.Println("时间到03")

}