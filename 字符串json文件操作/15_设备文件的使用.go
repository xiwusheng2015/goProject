package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("are you ok")
	os.Stdout.WriteString("hello nihao")

	var a int
	fmt.Println("请输入数字: ")
	fmt.Scan(&a)
	fmt.Println("a = ",a)
}
