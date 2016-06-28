package chapter3

import "fmt"

func reset(value int) {
	value = 0
}

func resetPointer(value *int) {
	*value = 0
}

func Pointers() {
	value := 100

	fmt.Print("initial:", value)

	reset(value)

	fmt.Print(" reset:", value)

	resetPointer(&value)

	fmt.Print(" pointer:", value)

	fmt.Println(" address:", &value)
}
