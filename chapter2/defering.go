package chapter2

import "fmt"

// Simple use of defering.
func Defering() {
	defer fmt.Println("World")

	fmt.Print("Hello")
}