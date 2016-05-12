package chapter9

import (
	"net"
	"fmt"
	"bufio"
	"time"
)

func createServer() func (stop chan struct{}) {
	fmt.Println("Starting server")
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	handleConnection := func(conn net.Conn, stopping chan struct{}) {
		for {
			select {
			case <- stopping:
				defer conn.Close()
				break
			default:
				message, _ := bufio.NewReader(conn).ReadString('\n')
				fmt.Println(conn.RemoteAddr().String(), "client :", message)
				conn.Write([]byte("ack" + "\n"))
			}
		}
	}

	return func (stop chan struct{}) {
		for {
			conn, err := ln.Accept()

			if err != nil {
				panic(err)
			}

			select {
			case <- stop:
				defer ln.Close()
				break
			default:
				go handleConnection(conn, stop)
			}
		}
	}
}

func createClient() func (stop chan struct{}) {
	fmt.Println("Connecting to server")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		panic(err)
	}

	return func (stop chan struct{}) {
		for i := 0; i < 10; i++ {
			fmt.Fprintf(conn, "hello\n")

			message, _ := bufio.NewReader(conn).ReadString('\n')

			fmt.Println("Server Ack: ", message)
		}

		defer conn.Close()

		close(stop)
	}
}

// Socket server example using closures, non-blocking go routines with stopping channels.
func Sockets() {
	stop := make(chan struct{})
	server := createServer()

	go server(stop)

	client := createClient()

	// waiting just in case things are slow
	time.Sleep(time.Second * 1)

	go client(stop)

	// stopping server and all go routines
	<-stop
}