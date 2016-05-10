package chapter1

import "fmt"

func Maps() {
	var m = make(map[string]int)

	m["a"] = 43
	m["b"] = 56

	fmt.Println(m)

	delete(m, "b")

	fmt.Println(m)

	value, prs := m["a"]

	fmt.Println(value, "and", prs)

	nothing, test := m["hello"]

	fmt.Println(nothing, "and", test)	
}