package chapter2

import "fmt"

// Demonstrating how to instanciate an inline function and executing it.
func InLineFunctions() {
	printHello := func () {
		fmt.Println("Hello")
	}

	printHello()
}