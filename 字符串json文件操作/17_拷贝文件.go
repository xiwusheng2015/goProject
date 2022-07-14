package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//获取命令行参数
	list := os.Args
	if len(list) != 3 {
		fmt.Println("usage : xxx srcFile dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]

	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件一样")
		return
	}

	// 只读方式打开源文件
	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}
	// 新建目的文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	//函数结束后关闭文件
	defer sF.Close()
	defer dF.Close()
	// 核心处理,从源文件读取内容，往目的文件写，读多少写多少
	buf := make([]byte, 1024*4)
	for {
		n, err := sF.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
			return
		}
		// 写入目标文件,读多少写多少
		dF.Write(buf[:n])
	}
}
