package main;

import (
	"github.com/dmportella/golang-tutorial/chapter0"
	"github.com/dmportella/golang-tutorial/chapter1"
	"github.com/dmportella/golang-tutorial/chapter2"
	"github.com/dmportella/golang-tutorial/chapter3"
	"github.com/dmportella/golang-tutorial/chapter4"
	"github.com/dmportella/golang-tutorial/chapter6"
)

func main() {
	chapter0.Printing()
	chapter0.Variables()
	chapter0.ControlFlow()
	chapter0.IfElse()
	chapter0.Switches()

	chapter1.Arrays()
	chapter1.Slices()
	chapter1.Maps()
	chapter1.Range()

	chapter2.Functions()
	chapter2.InLineFunctions()
	chapter2.MultiReturn()
	chapter2.Variadric()
	chapter2.Closure()
	chapter2.Recursion()
	chapter2.Defering()
	chapter2.StackingDefers()

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
	chapter4.Directions()

	chapter6.Selecting()
	chapter6.Timeouts()
	chapter6.Ranging()
	chapter6.NonBlocking()
}