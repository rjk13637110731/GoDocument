package main

import "fmt"

func main() {
	var a int = 21
	var c int
	c = a
	fmt.Printf("第1行- c = %d\n", c)
	c += a
	fmt.Printf("第2行- c = %d\n", c)
	c -= a
	fmt.Printf("第3行- c = %d\n", c)
	c *= a
	fmt.Printf("第4行- c = %d\n", c)
	c /= a
	fmt.Printf("第5行- c = %d\n", c)

	c = 200
	c <<= 2
	fmt.Printf("第6行- c = %d\n", c)
	c >>= 2
	fmt.Printf("第7行- c = %d\n", c)
	c &= 2
	fmt.Printf("第8行- c = %d\n", c)
	c ^= a
	fmt.Printf("第9行- c = %d\n", c)
	c |= a
	fmt.Printf("第10行- c = %d\n", c)
}
