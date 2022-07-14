package main

import "fmt"

type Person struct {
	id   int
	name string
	sex  byte
}

type Student struct {
	Person
	age  int
	addr string
}

func (tmp *Person) PrintInfo() {
	fmt.Printf("id = %d,name = %s,sex = %c\n", tmp.id, tmp.name, tmp.sex)
}

func main() {
	s := Student{Person{1, "nike", 'c'}, 12, "chongqing"}

	s.PrintInfo()
	s.Person.PrintInfo()

}
