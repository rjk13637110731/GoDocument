package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "3.14"

	num, err := strconv.ParseFloat(str, 64)

	if err != nil {
		fmt.Println("转换错误为：", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为浮点数为：%f\n", str, num)
	}
}
