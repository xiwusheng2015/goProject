package main

import "fmt"

func main() {
	s := "abc"
	for i,item := range s { //支持 string/array/slice/map。
		fmt.Printf("%d,%c\n", i,item)
	}
	arr := [3]int16{12, 424, 4444}
	for j := range arr {
		fmt.Printf("arr traverse %d\n", arr[j])
	}
	for _, c := range s { // 忽略 index
		fmt.Printf("%c\n", c)
	}


}
