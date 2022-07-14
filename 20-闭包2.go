package main

import "fmt"

func main22() {
	i := 10
	str := "mike"
	fmt.Printf("外部：i = %d, str = %s\n", i, str)
	func() {
		i = 100
		str = "go"
		//内部：i = 100, str = go
		fmt.Printf("内部：i = %d, str = %s\n", i, str)
	}() //别忘了后面的(), ()的作用是，此处直接调用此匿名函数

	//外部：i = 100, str = go
	fmt.Printf("外部：i = %d, str = %s\n", i, str)
}

// squares返回一个匿名函数，func() int
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x*x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

