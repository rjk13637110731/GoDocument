package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误：", err)
	} else {
		fmt.Printf("字符串'%s'转换为整数为：%d\n", str, num)
	}
}
