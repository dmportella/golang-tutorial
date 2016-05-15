package chapter1

import (
	"fmt"
)

func VariableRedux() {
	// You can instanciate lots of variables in a group like this.
	var (
		isBoolean bool       = true
		intMax  uint64     = 1<<64 - 1
		complex complex128 = 2+3i
	)
	
	// standard format would use %T(%v) however that would require us to pass the variables twice to print correctly
	// instead we use the format index, %[1]T(%[1]v) which tells the formatter to use the first variable we pass in.
	const f = "%[1]T(%[1]v)\n"

	fmt.Printf(f, isBoolean)
	fmt.Printf(f, intMax)
	fmt.Printf(f, complex)
}