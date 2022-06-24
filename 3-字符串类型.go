package main

import "fmt"

func main (){
	var str1 string
	str1 = "abc"
	fmt.Println("str1 = ",str1)

	str2 := "mike"
	fmt.Printf("str2 type = %T\n",str2)

	fmt.Println("len(str2) is ",len(str2))

	fmt.Println("============")

	// 字符串和字符区别
	dog := "hello go"
	fmt.Printf("str first id %c, second is %c\n", dog[0],dog[1])

}
