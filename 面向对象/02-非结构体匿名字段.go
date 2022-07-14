package main

import "fmt"

type mystr string

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	int
	mystr
}

func main() {

	s := Student{Person{"mike", 'm', 33}, 666, "hehe"}

	fmt.Printf("s = %+v\n", s)

	s.age = 22
	s.name = "c++"
	s.int = 923
	s.mystr = "my"

	fmt.Println(s.age, s.name, s.int, s.mystr)

	s.Person = Person{name: "xiaoming"}
	fmt.Println(s.Person, s.int, s.mystr)

}
