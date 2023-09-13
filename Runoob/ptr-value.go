package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	fmt.Printf("交换之前a = %d\n", a)
	fmt.Printf("交换之前b = %d\n", b)

	swap(&a, &b)
	fmt.Printf("交换之后a = %d\n", a)
	fmt.Printf("交换之后b = %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
