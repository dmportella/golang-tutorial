package chapter1

import (
	"fmt"
)

// Examples of arrays in golang
func Arrays() {

	// arrays have always a size which can be nil or more.
	var a [5]int

	// you can address each element of the array by its given index
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4

	fmt.Println(a)

	// you can declare and instanciate an array inline.
	b:= [5]int{1,2,3,4,5}

	fmt.Println(b)


	// simple loop adding all the values of the array
	sum:= 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	fmt.Println(sum)
}
