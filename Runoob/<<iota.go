package main

import "fmt"

func main() {
	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
	fmt.Println("l = ", l)
}
