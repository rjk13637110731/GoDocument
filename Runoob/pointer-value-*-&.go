package main

import "fmt"

func main() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("a变量的地址是：%x\n", &a)
	fmt.Printf("ip变量存储的指针地址：%x\n", ip)
	fmt.Printf("*ip 变量的值为：%d\n", *ip)
}
