package main

import "fmt"

func main () {
	var ch byte
	ch = 97
	fmt.Println("ch = ",ch)
	fmt.Printf("%c, %d\n",ch ,ch)
	ch = 'a'
	fmt.Printf("%c,%d\n",ch,ch)

	fmt.Printf("大写: %c, 小写：%c\n",'A','a')
	fmt.Printf("大写: %c, 转为小写：%c",'A','A'+32)

	fmt.Printf("hello go%c",'\n')
	fmt.Printf("hello,itcast")
}
