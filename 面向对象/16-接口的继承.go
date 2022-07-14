package main

import "fmt"

type Humaner interface {
	Sayhi()
}

type Personer interface {
	Humaner
	Sing(lrc string)
}

type Student struct {
	id   int
	name string
}

// Student 实现 Sayhi() 方法
func (tmp *Student) Sayhi() {
	fmt.Printf("Studentd[%d,%s]\n", tmp.id, tmp.name)

}

// Student 实现 Sing() 方法
func (tmp *Student) Sing(lrc string) {
	fmt.Printf("Student[%d,%s] Sing %s\n", tmp.id, tmp.name, lrc)
}

func main() {
	var p Personer
	s := Student{123, "gogogo"}
	p = &s
	p.Sayhi()
	p.Sing("hahaha")
}
