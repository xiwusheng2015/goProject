package main

import "fmt"

func MyFuncTT(args ...int) {
	fmt.Println("MyFunc01")
	for _, n := range args {
		fmt.Println(n)
	}
}

func MyFunc02(args ...int) {
	fmt.Println("MyFunc02")
	for _, n := range args { //遍历参数列表
		fmt.Println(n)
	}

}

func Test(args ...int) {
	MyFuncTT(args...)
	MyFunc02(args[1:]...)
}

func main() {
	//Test01(33,44)
	Test(1, 2, 3)
}
