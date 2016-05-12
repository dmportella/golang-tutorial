package chapter2

import "fmt"

func counter() func() int {
	count:=0
	return func() int {
		count++
		return count
	}
}

// Demonstrating clouses at work.
func Closure() {
	myCounter := counter()

	fmt.Println("#", myCounter())
	fmt.Println("#", myCounter())
	fmt.Println("#", myCounter())

	myCounter2 := counter()

	fmt.Println("#", myCounter2())
}