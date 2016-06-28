package chapter4

import "fmt"

func Unbuffered() {
	messages := make(chan string, 2)

	messages <- "Hello"
	messages <- "World"

	fmt.Print(<-messages)
	fmt.Println(<-messages)
}
