package chapter0

import "fmt"

func IfElse() {
	if 4 % 2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if 4 % 2 == 0 {
		fmt.Println("4 is divisibly by 2")
	}

	if num := 9; num < 10 {
		fmt.Println(num, "is less the 10")
	} 
}