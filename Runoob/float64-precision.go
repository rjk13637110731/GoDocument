package main

import "fmt"

func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}

func main() {
	a := 1690
	b := 1700
	c := a * b
	fmt.Println(c)
	fmt.Println(float64(c) / 1000000)
}
