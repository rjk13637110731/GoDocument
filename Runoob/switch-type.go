package main

import "fmt"

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型是：%T", i)
	case int:
		fmt.Printf("x 的类型是 int")
	case float64:
		fmt.Printf("x 的类型是 float64")
	case func(int) float64:
		fmt.Printf("x 的类型是 func(int)")
	case bool, string:
		fmt.Printf("x 的类型是 bool or string")
	default:
		fmt.Printf("x 的类型是Unknow")
	}
}
