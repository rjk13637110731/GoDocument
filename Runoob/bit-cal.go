package main

import "fmt"

func main() {
	var a uint = 60
	var b uint = 13
	var c uint = 0

	c = a & b
	fmt.Printf("第1行 -c的值为：%d\n", c)
	c = a | b
	fmt.Printf("第2行 -c的值为：%d\n", c)
	c = a ^ b
	fmt.Printf("第3行 -c的值为：%d\n", c)
	c = a << 2
	fmt.Printf("第4行 -c的值为：%d\n", c)
	c = a >> 2
	fmt.Printf("第5行 -c的值为：%d\n", c)
}
