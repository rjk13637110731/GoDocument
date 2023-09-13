package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	if a == 100 {
		if b == 200 {
			fmt.Printf("a = 100，b = 200\n")
		}
	}
	fmt.Printf("a的值为： %d\n", a)
	fmt.Printf("b的值为： %d\n", b)
}
