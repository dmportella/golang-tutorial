package chapter1

import (
	"fmt"
)

func Arrays() {
	var a [5]int

	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4

	fmt.Println(a)

	b:= [5]int{1,2,3,4,5}

	fmt.Println(b)

	sum:= 0
	for _, bint := range b {
		sum += bint
	}

	fmt.Println(sum)
}
