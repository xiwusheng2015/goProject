package main

import "fmt"

func main() {
	const n int = 10
	var a [n]int
	var b [5]int
	fmt.Printf("a is %v,b is %v\n", a, b)
	fmt.Printf("&a is %v,&b is %v", &a, &b)

	fmt.Println("===========")
	//var c [10]int
	for i := 0; i < 10; i++ {
		a[i] = i + 1
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
	fmt.Printf("a is %v,b is %v\n", a, b)
	for i, v := range a{
		fmt.Printf("range: a[%d] = %d\n",i,v)
	}

	// len和cap函数
	fmt.Printf("a's len is %d, a's cap is %d",len(a),cap(a))

}
