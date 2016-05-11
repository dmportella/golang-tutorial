package chapter4

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func Directions() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "ping pongs message")
	
	pong(pings, pongs)

	fmt.Println(<-pongs)
}