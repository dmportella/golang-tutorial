package chapter1

import "fmt"

func Range() {
	numbers:= [5]int{1, 2, 3, 4 ,5}

	for index, value := range numbers {
		fmt.Println("index:", index, "value:", value)
	}
}