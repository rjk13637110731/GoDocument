package main

import "fmt"

func main() {
	var a int = 20
	var ip *int
	var ptr *int
	ip = &a
	fmt.Printf("a变量的地址是：%x\n", &a)
	fmt.Printf("ip变量存储的指针地址：%x\n", ip)
	fmt.Printf("*ip 变量的值为：%d\n", *ip)
	fmt.Printf("ptr的值为：%x\n", ptr)
}
