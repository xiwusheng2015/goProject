package main

import "fmt"

type MyInt int

func (a MyInt) Add(b MyInt) MyInt {
	return a + b

}

func Add(a, b MyInt) MyInt {
	return a + b
}

func main() {
	var a MyInt = 1
	var b MyInt = 1

	fmt.Println("a.Add(b) = ", a.Add(b))

	fmt.Println("Add(a,b) = ", Add(a, b))
}
