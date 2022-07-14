package main

import "fmt"

func main(){
	var m1 map[int]string = map[int]string{1:"mike",2:"yoyo"}
	fmt.Println(m1)

	//2、自动推导类型 :=
	m2 := map[int]string{1: "mike", 2: "yoyo"}
	fmt.Println(m2)

}
