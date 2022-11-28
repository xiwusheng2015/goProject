package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/index.html", HandConn)

	// 监听绑定
	http.ListenAndServe(":8000", nil)
}

func HandConn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Mehtod = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)
	fmt.Println("有人访问了一下", r.RemoteAddr)

	w.Write([]byte("nihao ddd我是的发生的范德萨范德萨范德萨范德萨佛挡杀佛"))

}
