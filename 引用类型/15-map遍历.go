package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike", 2: "yoyo"}
	for k, v := range m1 {
		fmt.Printf("%d -----> %s\n", k, v)
	}

	//迭代遍历2，第一个返回值是key，第二个返回值是value（可省略）
	for k := range m1 {
		fmt.Printf("%d ----> %s\n", k, m1[k])
		//1 ----> mike
		//2 ----> yoyo
	}

	//判断某个key所对应的value是否存在, 第一个返回值是value(如果存在的话)
	value, ok := m1[6]
	fmt.Println("value = ", value, ", ok = ", ok) //value =  mike , ok =  true
	value2, ok2 := m1[3]
	fmt.Println("value2 = ", value2, ", ok2 = ", ok2) //value2 =   , ok2 =  false

	// 删除
	delete(m1, 9)
	for k, v := range m1 {
		fmt.Printf("%d ----> %s\n", k, v)

	}

}
