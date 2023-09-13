package main

import "fmt"

var g int = 20

func main() {
	var a, b int
	a = 10
	b = 20
	var g int = a + b

	fmt.Printf("结果是：a = %d, b = %d and g = %d\n", a, b, g)
}
