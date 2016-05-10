package main;

import (
	"tutorial/chapter0"
	"tutorial/chapter1"
	"tutorial/chapter2"
	"tutorial/chapter3"
	"tutorial/chapter4"
)

func main() {
	chapter0.Printing()

	chapter1.Arrays()
	chapter1.Slices()
	chapter1.Maps()
	chapter1.Range()

	chapter2.Functions()
	chapter2.MultiReturn()
	chapter2.Variadric()
	chapter2.Closure()
	chapter2.Recursion()
	chapter2.Defering()

	chapter3.Pointers()
	chapter3.Structures()
	chapter3.Methods()
	chapter3.Interfaces()
	chapter3.Errors()

	chapter4.GoSimple()
	chapter4.Waiting()
	chapter4.Channels()
	chapter4.Unbuffered()
	chapter4.Synchronising()
}