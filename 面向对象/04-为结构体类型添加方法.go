package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (tmp Person) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

func (p *Person) SetInfo(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	(*p).age = age
}

func main() {
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	var p2 Person
	(&p2).SetInfo("xioaming", 'f', 33)
	p2.PrintInfo()
}
