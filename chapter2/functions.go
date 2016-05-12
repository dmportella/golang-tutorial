package chapter2

import "fmt"

func addition(a int, b int) int {
	return a + b
}

// Simple functions at work.
func Functions() {
	result:= addition(10, 10)

	fmt.Println("10 + 10 =", result)
}