package chapter6

import "fmt"
import "time"

func red(text string) string {
	return fmt.Sprintf("\x1b[41m\x1b[93m%s\x1b[0m", text)
}

func green(text string) string {
	return fmt.Sprintf("\x1b[42m%s\x1b[0m", text)
}

// Making use of time.After to timeout go routines this example
// also demonstrate the use of cancelling a go routine that is waiting
// on a another go routine. This example also demonstrate non-blocking
// time outs.
func Timeouts() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	// channel of struct is a channel of nothing which can only be closed to signal a go routine.
	fin := make(chan struct{})

	// cancellable go routine
	go func(msg chan<- string, finish chan struct{}) {
		select {
		case <-finish:
			fmt.Println(red("Channel 1 Closed!!!"))
			return
		case <-time.After(time.Second * 1):
			fmt.Println(green("DONE STUFF"))
			c1 <- "ping"
		}

	}(c1, fin)

	go func() {
		time.Sleep(time.Millisecond * 800)
		fmt.Println(green("channel 2 go routine"))
		c2 <- "pong"
	}()

	select {
	case responce := <-c1:
		fmt.Println(responce)
	case responce := <-c2:
		// once we receive this we will tell the other go routine to stop by signaling the fin channel with a close
		fmt.Println(responce)
		close(fin)
	case <-time.After(time.Second * 2):
		fmt.Println(red("timeout"))
	}

	time.Sleep(time.Second * 2)
}
