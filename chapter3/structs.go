package chapter3

import "fmt"

type person struct {
	name string
	age  int
}

func Structures() {
	bob := person{"bob", 20}

	fmt.Println("name: ", bob.name, "age: ", bob.age)

	fmt.Println(person{name: "Neo", age: 30})

	fmt.Println(person{name: "Morpheous"})

	fmt.Println(&person{name: "Trinity", age: 31})

	// pointers

	human := person{name: "neo", age: 30}

	fmt.Println(human.name)

	theOne := &human
	poser := human

	fmt.Println(theOne.age)

	theOne.age = 34

	human.name = "the one"

	fmt.Println(theOne, human, poser)
}
