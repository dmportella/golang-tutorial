package chapter1

import (
	"fmt"
)

func Slices() {
	slice := make([]string, 3)

	fmt.Println(slice)

	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"

	fmt.Println(slice)

	slice = append(slice, "d")
	slice = append(slice, "e", "f")

	fmt.Println(slice)

	cpy:= make([]string, len(slice))

	copy(cpy, slice)

	fmt.Println(cpy)

	cpy = append(cpy[:2], cpy[3:]...)

	fmt.Println(cpy)
}