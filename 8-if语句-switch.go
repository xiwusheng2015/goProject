package main

import "fmt"

func main() {
	switch s3 := 90; { //只有初始化语句，没有条件
	case s3 >= 90: //这里写判断语句
		fmt.Println("优秀")
	case s3 >= 80:
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}

}
