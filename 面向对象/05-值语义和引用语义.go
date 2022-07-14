package main

import "fmt"

type Person struct {
	id int
	name string

	age int
}

// 值传递
func (tmp Person) setInfoValue(id int, name string, age int){
	fmt.Printf("&tmp = %p\n",&tmp)
	tmp.id = id
	tmp.name = name
	tmp.age = age
	fmt.Println("tmp = ",tmp)
}

// 引用传递
func (p *Person)setInfoPointer(id int, name string, age int)  {
	fmt.Printf("&p = %p\n",p)

	p.id = id
	p.name = name
	p.age = age
	fmt.Println("p = ",p)

}


func main() {

	var p1 Person = Person{1,"go",11}
	fmt.Printf("&p1= %p\n",&p1)

	//p1.setInfoValue(2,"c++",22)
	(&p1).setInfoPointer(3,"gg",3)
	fmt.Println("p1 = ",p1)

}
