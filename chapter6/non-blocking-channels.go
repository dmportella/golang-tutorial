package chapter6

import "fmt"

// Demonstrating how to fire a go routing with out block the main thread
// by using the select statement with a default case. The channel will not signal
// so the select will just fall into default and continue to run.
func NonBlocking() {
	channel := make(chan string)

	select {
	case message := <- channel:
		fmt.Println("message received", message)
	default:
		fmt.Println("main thread can continue")
	}

	message := "hello world"
	select {
		case channel <- message:
			fmt.Println("message sent", message)
		default:
			fmt.Println("main thread can continue")
	}
}