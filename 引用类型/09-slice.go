package main

import "fmt"

func main() {

	//var s1 [4]int
	//s2 := []int{}
	//fmt.Printf("%v,%T,%v,%T\n", s1,s1,s2, s2)
	//
	//var s3 []int = make([]int,0)
	//s4 := make([]int,0,0)
	//s5 := []int{1,2,3}
	//fmt.Printf("%v,%v,%v", s3,s4,s5)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := s[2:5]       //[2 3 4]
	s1[2] = 100        //修改切片某个元素改变底层数组
	fmt.Println(s1, s) //[2 3 100] [0 1 2 3 100 5 6 7 8 9]

	s2 := s1[2:6] // 新切片依旧指向原底层数组 [100 5 6 7]
	s2[3] = 200
	fmt.Println(s2) //[100 5 6 200]

	fmt.Println(s) //[0 1 2 3 100 5 6 200 8 9]



}
