package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("a = %d\n ", a)
	fmt.Printf("&a = %v \n ", &a)

	var p *int
	p = &a
	fmt.Printf("p = %v,&a = %v\n", p, &a)
	*p = 666
	fmt.Printf("p = %v,&a = %v, a = %v\n", p, &a, a)
}
