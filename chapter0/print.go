package chapter0

import "fmt"

// Examples of using fmt print methods
func Printing() {
	fmt.Print("Hellow world")

	fmt.Println("!!!!")

	fmt.Print("1.2.3.4.5")

	fmt.Println(".6")

	fmt.Println("hello world")

	fmt.Println("\x1b[41m\x1b[93mBOOM!!!\x1b[0m")

	fmt.Println(fmt.Sprintf("$ %0[3]*.[2]*[1]f", 155.445603, 2, 9))

	s:= ""
	i:= 0

	fmt.Sscanf(" 12 34 567 ", "%5s%d", &s, &i)

	fmt.Println(s, i)
}