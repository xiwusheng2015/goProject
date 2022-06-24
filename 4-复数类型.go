package main

import "fmt"

func main() {
	var t complex64
	t = 2.1  + 3.14i
	fmt.Println("t is ", t)

	t2 := 2.2 + 4.4i
	fmt.Printf("t2 type is %T\n",t2)

	// 分别获取实部和虚部
	fmt.Println("实部是",real(t2),"虚部是",imag(t2))
}
