package main

import (
	"encoding/json"
	"fmt"
)

// 成员变量名首字母必须大写
type IT struct {
	// Company string
	// Subject []string
	// Isok    bool
	// Price   float64
	Company string   `json:"Cub"`       // 此字段不会输出到屏幕
	Subject []string `json:"SSubject"` // 二次编码
	Isok    bool     `json:"is_ok"` // 类型改为 string 类型
	Price   float64  `json:",string"` // 类型改为 string 类型,冒号和引号之间不能有空格
}

func main() {
	s := IT{"com", []string{"go", "ccc"}, true, 11.11}
	buf, err := json.Marshal(s)
	//buf ,err := json.MarshalIndent(s,"*","	")

	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
}
