package chapter2

import "fmt"

// Advance use of defer, stacking.
func StackingDefers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
