package chapter4

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// Demonstrating how to block the main thread while a go routine is executing using channels.
func Synchronising() {
	done := make(chan bool, 1)

	go worker(done)

	<-done
}