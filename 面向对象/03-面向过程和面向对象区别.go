package main

import "fmt"

func Add01(a, b int) int {
	return a + b
}

// 面向对象： 方法： 给某个类型绑定一个函数
type long int

func (o1 long) Add02(o2 long) long {
	return o1 + o2
}

func main() {
	//var result int
	result := Add01(1, 2)

	fmt.Println("result is :", result)

	var a long = 4
	r := a.Add02(4)
	fmt.Println("r = ", r)

}
