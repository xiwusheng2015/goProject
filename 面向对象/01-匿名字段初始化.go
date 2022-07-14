package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id   int
	addr string
}

func main() {

	var s1 Student = Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Printf("s1 is %+v\n", s1)

	s3 := Student{id: 3}
	fmt.Printf("s3 is %+v\n", s3)

	s4 := Student{Person: Person{name: "mike"}, id: 99}
	fmt.Printf("s4 is %+v\n", s4)


}
