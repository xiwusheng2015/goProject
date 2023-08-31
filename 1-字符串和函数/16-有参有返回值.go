package main

import "fmt"

func MinAndMax(num1 int, num2 int) (min int, max int) {
	if num1 > num2 {
		min, max = num2, num1
	} else {
		min, max = num1, num2
	}
	return
}

func main() {
	min, max := MinAndMax(33, 44)
	fmt.Printf("min is %d, max is %d\n", min, max)
}
