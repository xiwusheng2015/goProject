package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 监听
	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	defer listener.Close()

	// 接收多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		// 新建一个协程处理
		go HandleConn(conn)
	}

}

func HandleConn(conn net.Conn) {
	// 函数结束时 ，自动关闭conn
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect successful")
	buf := make([]byte, 2048)

	// 读取
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println("len = ", len(string(buf[:n])))

		// 输出用户的地址
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
		if "exit" == string(buf[:n-2]) {
			fmt.Println(addr, " exited")
			return
		}

		// 把数据转换为大写，再发送给用户
		_, err1 := conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
		if err1 != nil {
			fmt.Println("write err", err.Error())
			return
		}
	}
}
