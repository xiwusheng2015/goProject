package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func (tmp Person) setInfoValue() {
	fmt.Println("setInfoValue")
}

func (tmp *Person) setInfoPointer() {
	fmt.Println("setInfoPointer")
}

func main() {
	var p *Person = &Person{1, "go", 22}
	(*p).setInfoPointer()
	(p).setInfoValue()
}
