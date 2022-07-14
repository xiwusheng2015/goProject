package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p %v \n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("SetPointerValue: %p %v \n", p, p)
}

func main() {
	p := Person{1, "java6", 11}
	fmt.Printf("&p =%p\n ", &p)
	p.SetInfoValue()

	// 方法值
	pValue := p.SetInfoValue
	pValue()

	pPointer := p.SetInfoPointer
	pPointer()

}
