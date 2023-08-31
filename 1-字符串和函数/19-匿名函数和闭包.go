package main

import "fmt"

func main() {

	i := 0
	str := "mike"
	// 方式一
	f1 := func() {
		//引用到函数外的变量
		fmt.Printf("方式1：i = %d, str = %s\n", i, str)
	}
	f1()
	//方式1的另一种方式
	type FuncType func() //声明函数类型, 无参无返回值
	var f2 FuncType = f1
	f2()

	// 方式二
	var f3 FuncType = func() {
		fmt.Printf("方式2：i = %d, str = %s\n", i, str)
	}
	f3() //函数调用

	// 方式三
	func() {
		fmt.Printf("方式3：i = %d, str = %s\n", i, str)
	}()

	// 方式四
	v := func(a, b int) (result int) {
		result = a + b
		return
	}(1, 2)

	fmt.Println("v = ", v)

}
