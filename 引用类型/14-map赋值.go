package main

import "fmt"

func main(){
	m1 := map[int]string{1: "mike", 2: "yoyo"}
	m1[1] = "xxx"   //修改
	m1[3] = "lily"  //追加， go底层会自动为map分配空间
	m1[999] = "999"
	fmt.Println(m1) //map[1:xxx 2:yoyo 3:lily]

	m2 := make(map[int]string, 2) //创建map
	m2[0] = "aaa"
	m2[1] = "bbb"
	m2[2] = "bbb"
	m2[3] = "bbb"
	m2[4] = "bbb"

	fmt.Println(m2)           //map[0:aaa 1:bbb]
	fmt.Println(m2[0], m2[1]) //aaa bbb

}
