package chapter9

import (
	"net"
	"fmt"
	"bufio"
	"time"
)

// Simple closure that creates a listening port (opening it) and 
// returning a anonymous function with the go routine loop for 
// accepting client connections. This server can be closed with
// the stop channel.
func createServer() func (stop chan struct{}) {
	fmt.Println("Starting server")
	ln, err := net.Listen("tcp", ":44646")

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

// Simple client closure that dial to the server and returns a anonymous
// function that handles the messaging with the server, sending 10 hellos
// and receiving the acknowlegements from the server. At the signaling
// for the closure of the client and server.
func createClient() func (stop chan struct{}) {
	fmt.Println("Connecting to server")
	conn, err := net.Dial("tcp", "127.0.0.1:44646")

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

	// When the client closes the stop channel the maint thread will get the signal and continue to execute.
	<-stop
}