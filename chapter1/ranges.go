package chapter1

import "fmt"

// ranges can bu used to iterate over arrays, maps, slices and channels.
// we will discuss channels at a later chapter.
func Range() {
	numbers:= [5]int{1, 2, 3, 4 ,5}

	// simple use of a range over our numbers
	for index, value := range numbers {
		fmt.Println("index:", index, "value:", value)
	}
}