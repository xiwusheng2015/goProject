package main

import "fmt"

func main() {
	test(3)
}

func test(i int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:")
			fmt.Println(err)
		}
	}()
	var a [10]int
	a[i] = 123
}
