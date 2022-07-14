package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "./demo.txt"
	//writeFile(path)
	//readFile(path)
	readFileLine(path)
}

/**
按行读取
*/
func readFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// 新建一个缓冲区，内容先放入缓冲区
	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("buf = #%s#\n", string(buf))
	}
}

func readFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 102*2) // 创建一个2k切片
	// n表示从文件读取的内容的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { // 文件出错，同时没有到结尾
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println("buf = \n", string(buf[:n]))
}

func writeFile(path string) {
	//打开文件,新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println(f)
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		// buf里面是i=0
		buf = fmt.Sprintf("i = %d\n", i)
		f.WriteString(buf)
	}

}
