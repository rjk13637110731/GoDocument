package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Printf("1、case false\n")
		fallthrough
	case true:
		fmt.Printf("2、case true\n")
		fallthrough
	case false:
		fmt.Printf("3、case false\n")
		fallthrough
	case true:
		fmt.Printf("4、case true\n")
	case false:
		fmt.Printf("5、case false\n")
		fallthrough
	default:
		fmt.Printf("6、默认case\n")
	}
}
