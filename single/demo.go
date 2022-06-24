
package main

import "fmt"

func main() {
	var a int
	fmt.Printf("input a please: ")
	fmt.Scanf("%d\n", &a)
	fmt.Println("a is ", a)
	var b float32
	fmt.Scan(&b)
	fmt.Println("b is ", b)
}

