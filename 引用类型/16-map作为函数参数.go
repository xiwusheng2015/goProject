package main

import "fmt"

func DeleteMap(m map[int]string, key int) {
	delete(m, key)

	for k, v := range m {
		fmt.Printf("in func: len(m)= %d, %d ------> %s\n", len(m), k, v)
	}
}

func main() {
	m := map[int]string{1: "make", 2: "youyou", 3: "lily"}

	DeleteMap(m, 9)

	for k, v := range m {
		fmt.Printf("len(m) = %d, %d ----> %s\n", len(m), k, v)
	}
}
