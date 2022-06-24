package main

import "fmt"

func test() (a, b,c int) {
	return 1,2,3
}

func main() {
	var (
		f1 float32

	)
	f1 = 2.3232323
	fmt.Println("f1 = ",f1)
	fmt.Printf("f1 type is %T\n", f1)
	f2 := 3.2323
	fmt.Printf("f2 type is %T\n", f2)

}