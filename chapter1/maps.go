package chapter1

import "fmt"

// Example of maps the dictionaries of golang.
// Maps are not thread safe we will discuss this at a later chapter.
func Maps() {
	// maps use the builtin function make to instanciate it.
	// you give it a index type or key type
	// and you give it a value type.
	var m = make(map[string]int)

	// you use the index type to set the values of the map
	m["a"] = 43
	m["b"] = 56

	// you can use the builtin len function to get the length of a map
	fmt.Println(m, len(m))

	// use the builtin delete function to delete an entry on a map 
	delete(m, "b")

	fmt.Println(m)

	// if you address a key that doesnt exist the value will be the zero value of the value type.
	n := m["nothing"]

	fmt.Println(m, n)

	// if you want to see if an index didnt exist you can use the multi return value
	// where 'value' is the value of the element and 'ok' is if it exists or not (boolean eg true or false)
	value, ok := m["a"]

	fmt.Println(value, "and", ok)

	nothing, test := m["hello"]

	fmt.Println(nothing, "and", test)	
}