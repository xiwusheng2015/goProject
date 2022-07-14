package main

import "fmt"

func main() {
	var p *int
	p = nil
	fmt.Println("p = ",p)

	var a int
	p = &a
	*p = 666
	fmt.Println("a = ",a)
}
