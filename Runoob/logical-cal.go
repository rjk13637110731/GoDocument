package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false

	if a && b {
		fmt.Printf("1.True\n")
	}
	if a || b {
		fmt.Printf("2.True\n")
	}

	a = false
	b = true
	if a && b {
		fmt.Printf("3.True\n")
	} else {
		fmt.Printf("3.False\n")
	}
	if !(a && b) {
		fmt.Printf("4.True\n")
	}
}
