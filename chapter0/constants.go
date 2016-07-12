package chapter0

import "fmt"

// constants types dont need to have their types explicitly set
const answer = 42

// Constants can be instanciated in a group
const (
	HELLO = 0
	WORLD = 1
)

// iota is a self assigning const that increments itself every time is used
const (
	zero = iota
	one
	two
)

// iota counter are unique per constant group
const (
	reseted = iota
	name    = "hello"
)

// Examples of using constants
// Show how each setup works and ues of iota
func Constants() {
	const f = "%[1]T(%[1]v)\n\r"

	// type is inferred when read
	fmt.Printf(f, answer)

	// casting is also possible
	fmt.Printf(f, int64(answer))

	fmt.Printf(f, HELLO)
	fmt.Printf(f, WORLD)

	// prints the values set by iota
	fmt.Printf(f, zero)
	fmt.Printf(f, one)
	fmt.Printf(f, two)

	// prints the reset iota and a string
	fmt.Printf(f, reseted)
	fmt.Printf(f, name)
}
