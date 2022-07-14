package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2}
	fmt.Println(a == b, b == c) //true false

	var d [3]int
	//e := [2]int{4,4}
	d = c
	fmt.Println(d) //[1 2 3]

}
