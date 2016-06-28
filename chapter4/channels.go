package chapter4

import "fmt"

func Channels() {
	messages := make(chan string)

	go func(message string) {
		messages <- message + " pong"
	}("ping")

	msg := <-messages

	fmt.Println(msg)
}
