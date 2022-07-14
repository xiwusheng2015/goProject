package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main22() {
	fmt.Println("")
	var s1 Student = Student{1, "mike", 'm', 18, "shanghai"}
	s2 := Student{2, "youyou", 'f', 22, "shenzhen"}

	//s3 := Student{2, "tom", 'm', 20} //err, too few values in struct initializer

	//2、指定初始化某个成员，没有初始化的成员为零值
	s4 := Student{id: 2, name: "lily"}

	fmt.Println(s1, s2, s4)

}

func main(){
	var s5 *Student = &Student{3,"xiaoming",'m',12,"beijing"}
	s6 := &Student{4,"roco",'m',44,"ddd"}

	fmt.Println(*s5,s6)
}
