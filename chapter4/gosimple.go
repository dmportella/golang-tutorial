package chapter4

import "fmt"
import "time"

func counting(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(1000 * time.Millisecond)
	}
}

// Simple example of go routines at work
func GoSimple() {
	counting("directly")

	go counting("async")

	go func(name string) {
		counting(name)
	}("anonymous")

	var input string

	fmt.Scanln(&input)

	fmt.Println("Done counting: ", input)
}