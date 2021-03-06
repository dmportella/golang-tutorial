package chapter0

import "fmt"

// Simple uses of for loops.
// Standards uses and infinite loops.
func ForLoop() {
	// you can instanciate variables outside of the for loop and only eval the expression in it.
	// this is also known as a while loop
	i := 0
	for i < 3 {
		fmt.Print(i)
		i = i + 1
	}

	fmt.Println()

	// standard for loop with val,  expression and increment
	for j := 0; j < 3; j++ {
		fmt.Print(j)
		j = j + 1
	}

	fmt.Println()

	// in golang there is no DO or WHILE statements you can achieve them by creating a empty for loop
	for {
		fmt.Println("infinite loop")
		break
	}
}
