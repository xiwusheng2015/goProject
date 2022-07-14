package main

import "fmt"

type Student struct {
	id   int
	name string
}

func main() {
	var i = make([]interface{}, 3)

	i[0] = 111
	i[1] = "hiell"
	i[2] = Student{11, "mike"}
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] 类型是int,值是%d\n", index, value)
		case string:
			fmt.Printf("x[%d] 类型是string，值是%s\n", index, value)
		case Student:
			fmt.Printf("x[%d]的类型是Student{}, 值id = %d, name = %s\n", index, value.id, value.name)
		}

	}
}
