package chapter2

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Simple example of recursion.
func Recursion() {
	fmt.Println(factorial(7))
}
