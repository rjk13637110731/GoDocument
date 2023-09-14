package main

import "fmt"

func sqrtRecursive(x, guess, preGuess, epsilon float64) float64 {
	if diff := guess*guess - x; diff < epsilon && -diff < epsilon {
		return guess
	}

	newGuess := (guess + x/guess) / 2
	if newGuess == preGuess {
		return guess
	}
	return sqrtRecursive(x, newGuess, guess, epsilon)
}

func sqrt(x float64) float64 {
	return sqrtRecursive(x, 1.0, 0.0, 1.0e-9)
}

func main() {
	x := 25.0
	result := sqrt(x)
	fmt.Printf("%.2f 的平方根为 %.6f\n", x, result)
}
