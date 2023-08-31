package main

import "fmt"

func Test01() int {
	i := 1
	sum := 0
	for i = 1; i <= 100; i++ {
		sum += i
	}
	return sum
}

func Test02(num int) int {
	if num == 1 {
		return 1
	}
	return num + Test02(num-1)
}

//通过递归实现1+2+3……+100
func Test03(num int) int {
	if num == 100 {
		return 100
	}

	return num + Test03(num+1) //函数调用本身
}

func main() {
	fmt.Println(Test01())
	fmt.Println(Test02(100))
	fmt.Println(Test03(1))
}
