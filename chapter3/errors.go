package chapter3

import "fmt"
import "errors"

func divide(a complex128, b complex128) (complex128, error) {
	if a == 0i || b == 0i {
		return 0, errors.New("Can not divide by zero!")
	}
	return a / b, nil
}

// Demonstrating the Error partner in golang.
func Errors() {
	var a, b complex128
	a, b = 100i, 3i

	if c, err := divide(a,b); err == nil {
		fmt.Printf("\x1b[42m%[1]g / %[2]g = %[3]g\x1b[0m\n\r", a, b, c)
	}

	var d complex128

	if c, err := divide(d,b); err == nil {
		fmt.Printf("%[1]g / %[2]g = %[3]g\n\r", a, b, c)
	} else {
		fmt.Println("\x1b[101m", err.Error(), "\x1b[0m")
	}
}