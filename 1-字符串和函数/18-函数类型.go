package main

import "fmt"

type FuncType func(int, int) int

//函数中有一个参数类型为函数类型：f FuncType
func Calc(a, b int, f FuncType) (result int) {
	result = f(a, b)
	return
}

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func main() {
	//函数调用，第三个参数为函数名字，此函数的参数，返回值必须和FuncType类型一致
	result := Calc(2, 3, Add)
	fmt.Println(result)

	var f FuncType = Minus
	fmt.Println("result is ", f(10, 2))
	fmt.Println("new minus is ", Calc(2, 3, f))
}
