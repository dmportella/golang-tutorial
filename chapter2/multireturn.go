package chapter2

import "fmt"

func addAndDivide(a int, b int) (int, int) {
	return (a + b), (a / b)
}

func MultiReturn() {
	add, div := addAndDivide(10, 2)

	fmt.Println("add:", add, "div:", div)

	_, b := addAndDivide(10, 6)

	fmt.Println("div:", b)
}