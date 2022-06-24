package main

import "fmt"

func main() {

	for i:=0;i<5;i++{
		if i==2{
			fmt.Println("i==2")
			//break
			continue
		}
		fmt.Println(i)
	}
}
