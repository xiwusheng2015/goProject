package main

import (
	"errors"
	"fmt"
)


func main() {
	var error1 error = fmt.Errorf("this is a error1")
	fmt.Printf("error1: %v\n", error1)

	error2 := fmt.Errorf("this is normal error2")
	fmt.Printf("error2: %s\n", error2)

	error3 := errors.New("this is new normal error3")
	fmt.Printf("error3: %s\n", error3)
}
