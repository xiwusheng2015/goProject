package main

import "fmt"

type Student struct {
	id   int
	name string
}

func main() {
	i := make([]interface{}, 3)
	i[0] = "hello"
	i[1] = 124
	i[2] = Student{2222, "mike"}

	for index, data := range i {
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型是int, 内容是%d\n", index, value)
		} else if value, ok := data.(string) ; ok== true{
			fmt.Printf("x[%d] 类型是string ，内容是%s\n",index, value)
		}else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为 Student{}, 内容为id = %d, name = %s\n", index, value.id, value.name)
		}
	}
}
