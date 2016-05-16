package chapter0

import "fmt"

func meaningOfLife() int {
	return 42
}

// Examples of If Else statements
func IfElse() {
	// normal if statement
	if 4 % 2 == 0 {
		fmt.Println("4 is divisibly by 2")
	}

	// an if statement with a else clause
	if 4 % 2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// in golang you can instanciate a variable and use it strait away in the if statement
	if num := 9; num < 10 {
		fmt.Println(num, "is less the 10")
	}

	// this also works with method calls
	// this is commonly used when i method with a multi return format returns a error
	// which you can use to evaluate if the method call was successful.
	if answer := meaningOfLife(); answer == 42 {
		fmt.Println("we found it!")
	}

	// a standard else if statement
	if num := 1; num == 0 {
		fmt.Println("num is zero")
	} else if num != 0 {
		fmt.Println("num is not zero")
	}
}