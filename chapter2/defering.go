package chapter2

import "fmt"

func Defering() {
	defer fmt.Println("World")

	fmt.Print("Hello")
}