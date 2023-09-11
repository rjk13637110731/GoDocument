package main

import "fmt"

var xx, yy int
var (
	aa int
	bb bool
)
var cc, dd int = 1, 2
var ee, fff = 123, "Hello"

func main() {
	var a string = "Runoob"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d int
	fmt.Println(d)

	var e bool
	fmt.Print(e)

	var i int
	var f float64
	var g bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, g, s)

	var h = true
	fmt.Println(h)

	ff := "Runoob"
	fmt.Println(ff)

	gg, hh := 123, "Hello"
	fmt.Println(xx, yy, aa, bb, cc, dd, ee, fff, gg, hh)
	fmt.Println(&xx)
}
