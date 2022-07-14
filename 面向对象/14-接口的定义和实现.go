package main

import "fmt"

type Humaner interface {
	say()
}

type Student struct {
	id   int
	name string
}

func (tmp *Student) say() {
	fmt.Printf("Student[%d,%s] say  \n", tmp.id, tmp.name)
}

type Teacher struct {
	sex  byte
	addr string
}

// Teacher实现 say() 方法
func (tmp Teacher) say() {
	fmt.Printf("Teacher[%c, %s] say \n", tmp.sex, tmp.addr)
}

type Mystr string

// Mystr 实现 say() 方法
func (tmp Mystr) say() {
	fmt.Printf("Mystr[%s] say\n", tmp)
}

func main() {
	var h Humaner
	s := Student{1, "mike"}
	h = &s
	h.say()
	fmt.Println("1  h = ", h)

	t := Teacher{'c', "niha"}
	h = t
	h.say()
	fmt.Println("2  h = ", h)

	var str Mystr = "me"
	h = str
	h.say()
}
