package chapter2

import "fmt"

func sum(nums ...int){
	fmt.Print(nums, " = ")

	total:= 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Simple use of '...'  variadric parameters.
func Variadric() {
	sum(1, 2, 3, 4, 5)

	nums:= []int{1,2,3,4,5,6}

	sum(nums...)
}