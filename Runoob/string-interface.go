package main

import (
	"fmt"
)

func main() {
	var i interface{} = "Hello World!"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("Conversion Failed")
	}
}
