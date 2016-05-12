package chapter6

import "fmt"
import "time"

func red(text string) string {
	return fmt.Sprintf("\x1b[41m\x1b[93m%s\x1b[0m", text)
}

func Timeouts() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	fin := make(chan struct{})

	go func(msg chan<- string, finish chan struct{}) {
		select {
		case <- finish:
			fmt.Println(red("Channel 1 Closed!!!"))
			return
		case <- time.After(time.Second * 1):
			fmt.Println("DONE STUFF")
			c1 <- "ping"
		}
		
	}(c1, fin)

	go func() {
		time.Sleep(time.Millisecond * 800)
		fmt.Println("channel 2 go routine")
		c2 <- "pong"
	}()

	select {
    case responce := <-c1:
        fmt.Println(responce)
    case responce := <-c2:
        fmt.Println(responce)
        close(fin)
    case <-time.After(time.Second * 2):
        fmt.Println("timeout")
    }
}