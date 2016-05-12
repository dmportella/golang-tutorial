package chapter6

import "fmt"

// receiving messages from a channel that is closed using a for range loop.
func Ranging() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3

	// if we dont close the range will throw an fatal error
	close(c)

	for v := range c{
		fmt.Println("value", v, "lenght", len(c))
	}

	// demonstrating how the channel looks like onc eis empty and closed
	v, ok := <- c

	fmt.Println("value", v, "lenght", len(c), "ok?", ok)

}