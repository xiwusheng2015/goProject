package main

import "fmt"

func main() {
	var i interface{}
	i = 1
	fmt.Println("i = ",i)

	i = "abc"
	fmt.Println("i = ",i)
}