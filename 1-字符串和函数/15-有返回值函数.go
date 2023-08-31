package main

import "fmt"

func Test01() int {
	return 222
}

//官方建议：最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差
func Test02() (value int) { //方式2, 给返回值命名
	value = 250
	return value
}

func Test03() (value int) { //方式3, 给返回值命名
	value = 250
	return
}

//多个返回值
func Demo1() (int, string) {
	return 250, "sb"
}

func Demo2() (a int, str string) {
	a = 250
	str = "sb"
	return
}

func main() {
	test02 := Test02()
	fmt.Println("result is ", test02)
	test03 := Test03()
	fmt.Println("result is ", test03)
	dd, ss := Demo1()
	fmt.Printf("first is %d, second is %s\n", dd, ss)
}
