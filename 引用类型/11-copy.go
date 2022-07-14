package main

import "fmt"

func main2() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := data[8:] //{8, 9}
	s2 := data[:5] //{0, 1, 2, 3, 4}
	copy(s2, s1)   // dst:s2, src:s1

	fmt.Println(s2)   //[8 9 2 3 4]
	fmt.Println(data) //[8 9 2 3 4 5 6 7 8 9]

}

func test(s []int) {
	s[0] = 999
	fmt.Println("test : ")
	for i, v := range s {
		fmt.Printf("s[%d] = %d\n", i, v)
	}
	fmt.Printf("\n")
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	test(slice)

	fmt.Println("main:")

	for i, v := range slice {
		fmt.Printf("slice[%d] = %d\n", i,v)

	}
}


