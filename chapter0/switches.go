package chapter0

import "fmt"

func Switches() {
	i := 2
	switch i {
	case 1:
		fmt.Println("never going to print")
	case 2, 3, 4, 5:
		fmt.Println("wow oke")
	default:
		fmt.Println("stop wasting time")
	}
}