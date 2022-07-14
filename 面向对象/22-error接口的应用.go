package main

import (
	"errors"
	"fmt"
)

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("0 不能是分母")
	} else {
		result = a / b
	}
	return
}

func main() {

	res, err := MyDiv(1, 0)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	} else {
		fmt.Printf("res: %d\n", res)
	}
}
