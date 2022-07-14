package main

import "fmt"

// 多个 defer 会先进后出
func main() {
	defer fmt.Println("aaaaaaaa")
	defer fmt.Println("bbbbbbbbb")
	// 调用一个函数，导致内存出问题

	defer fmt.Println("dddddd")
	test8(0)
	defer fmt.Println("ccccccc")
}

func test8(x int) {
	result := 100 / x
	fmt.Println("result = ", result)
}