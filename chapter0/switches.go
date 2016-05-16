package chapter0

import "fmt"

// Simple example of a switch statement
func Switches() {
	i := 2
	
	switch i {
	case 1:
		fmt.Println("never going to print")
		// case statements can be grouped
	case 2, 3, 4, 5: 
		fmt.Println("wow oke")
		// default is used when 'i' is not found in any case statement
	default:
		fmt.Println("stop wasting time")
	}
}