package main

import "fmt"

func Test01(v1 int, v2 int) { //方式1
	fmt.Printf("v1 = %d, v2 = %d\n", v1, v2)
}

func Test02(v1, v2 int) { //方式2, v1, v2都是int类型
	fmt.Printf("v1 = %d, v2 = %d\n", v1, v2)
}

//形如...type格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数
func Test(args ...int) {
	for _, n := range args { //遍历参数列表
		fmt.Println(n)
	}
}



func main() {
	Test(2,3,45,4)
	Test01(12,22)
	Test02(1,2)
}
