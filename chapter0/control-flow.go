package chapter0

import "fmt"

func ControlFlow() {
	i:= 0
	for i < 3 {
		fmt.Print(i)
		i = i + 1
	}
	fmt.Println()

	for j:= 0; j < 3; j++ {
		fmt.Print(j)
		j = j  + 1
	}
	fmt.Println()

	for {
		fmt.Println("infinite loop")
		break
	}
}