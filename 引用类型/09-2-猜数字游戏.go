package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNum(p *int) {
	// 设置种子
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}
	}
	*p = num
}

func GetNum(s []int, num int) {
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
}

func OnGame(randSlice []int) {
	var num int
	for {
		for {
			fmt.Printf("请输入一个4位数:")
			fmt.Scan(&num)
			if 999 < num && num < 10000 {
				break
			}
			fmt.Println("输入的数不符合要求")
		}
		fmt.Println("num = ", num)
		keySlice := make([]int, 4)
		GetNum(keySlice, num)
		fmt.Println("keySlice is : ", keySlice)

		// 开始对比是否猜对了
		f := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了", i+1)
			} else {
				fmt.Printf("第%d位猜对了", i+1)
				f++
			}
		}
		if f == 4 {
			fmt.Println("全部猜对了")
			break
		} else {
			fmt.Println("请继续猜")
		}
	}
}

func main() {
	var randNum int

	// 产生一个4位随机数
	CreateNum(&randNum)
	fmt.Println("randNum: ", randNum)
	// 切片
	randSlice := make([]int, 4)
	GetNum(randSlice, randNum)
	fmt.Println("randSlice is: ", randSlice)

	OnGame(randSlice)

}
