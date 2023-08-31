package main

import "fmt"

func main() {
	var ch byte = 97
	 a := int(ch)
	fmt.Println("a is ",a)
	 a++
	 a= 4
	 b := 4 << a
	fmt.Println("b is ",b)
}
