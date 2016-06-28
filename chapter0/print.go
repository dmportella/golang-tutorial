package chapter0

import "fmt"

// Examples of using fmt print methods
// The golang documentation is the ultimate reference for formating and the fmt library
// For the purpose of these examples i will show just a few but i encourage you to visit
// https://golang.org/pkg/fmt/
func Printing() {
	// Prints test to the current caret position on the screen
	fmt.Print("Hellow world")

	// prints text to the current caret position and adds a new line and carrege return
	fmt.Println("!!!!")

	fmt.Print("1.2.3.4.5")

	fmt.Println(".6")

	// simple example of using excape sequences to draw text in colour
	// check this site for more information about ascii excape sequences
	// http://misc.flogisoft.com/bash/tip_colors_and_formatting
	fmt.Println("\x1b[41m\x1b[93mBOOM!!!\x1b[0m")

	// Print'f' stands for print format which allows fine control over how variable are printed
	fmt.Printf("%T = %d", 123, 123)

	// same as the example above but we are using indexing on the format to avoid having to pass in the varaible twice
	fmt.Printf("%[1]T = %[1]d", 123)

	// this is a more complicated formated string useful for printing money
	fmt.Println(fmt.Sprintf("$ %0[3]*.[2]*[1]f", 155.445603, 2, 9))

	// Scanning is a feature that extracts data from a string using a given format.
	// in essense scans the string for something matching the format given
	s := ""
	i := 0

	fmt.Sscanf(" 12 34 567 ", "%5s%d", &s, &i)

	fmt.Println(s, i)
}
