package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) PrintInfo() {
	fmt.Println(p.name, p.sex, p.age)
}

func main() {
	p := Person{"mike", 'm', 19}
	p.PrintInfo()
}
