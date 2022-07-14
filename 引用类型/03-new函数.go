package main

import "fmt"

func main() {
	//a := 10
	var p *int
	//p = &a
	p = new(int)
	*p = 666
	fmt.Println("*p = ", *p, "p = ", p)

	q := new (int)
	*q = 777
	fmt.Println("*p = ", *q, "p = ", q)
}
