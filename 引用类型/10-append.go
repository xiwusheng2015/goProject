package main

import "fmt"

func main22() {
	var s1 []int //创建nil切换
	//s1 := make([]int, 0)
	s1 = append(s1, 1)       //追加1个元素
	s1 = append(s1, 2, 3)    //追加2个元素
	s1 = append(s1, 4, 5, 6) //追加3个元素
	fmt.Println(s1)          //[1 2 3 4 5 6]

	s2 := make([]int, 5)
	s2 = append(s2, 6)
	fmt.Println(s2) //[0 0 0 0 0 6]

	s3 := []int{1, 2, 3}
	s3 = append(s3, 4, 5)
	fmt.Println(s3)//[1 2 3 4 5]

}

func main() {
	s := make([]int, 0, 1)
	c := cap(s)
	for i := 0; i < 30; i++ {
		s = append(s, i)
		fmt.Println("current len is",len(s),"cap is",cap(s))
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
	/*
	   cap: 1 -> 2
	   cap: 2 -> 4
	   cap: 4 -> 8
	   cap: 8 -> 16
	   cap: 16 -> 32
	   cap: 32 -> 64
	*/
}

