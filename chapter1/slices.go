package chapter1

import (
	"fmt"
)

// Slices are the list type for golang they are a dynamic view into the inside of an array.
// they have better control over arrays, they dont store any data they are literally references to arrays.
func Slices() {
	// slices can be instanciated with the builtin function make
	// like an array they have a type the second param of make is for the capacity of the slice
	// slices can be appended too
	slice := make([]string, 3)

	// slices can be instanciated with the inline literal
	inline := []string{
		"hello",
		"world",
	}

	// slices values can be accessed by their index in the slice.
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"

	fmt.Println(slice, inline)

	// slices can be appended too by use of the append builtin function
	slice = append(slice, "d")
	slice = append(slice, "e", "f")

	fmt.Println(slice)

	cpy := make([]string, len(slice))

	// copy is another builtin function that can be used to copy the contents from one slice into another
	copy(cpy, slice)

	fmt.Println(cpy)

	// slices can be created from arrays
	array := []int{1, 2, 3}

	// using the indexing to create a slice over the array.
	// [start:end] notation instructs the runtime to copy the array from start index to the end index.
	// ommitting a value on start is equivalent to zero and omitting on the end is equivalent to len(array)
	sliced := array[0:]

	sliced[1] = 67

	// you will see that changing sliced you will be changing the array it references.
	fmt.Println(array, sliced)
}
