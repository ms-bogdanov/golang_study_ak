package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type client struct {
	conn net.Conn
	name string
	ch   chan string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Server started at localhost:8000")

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)

			continue
		}

		go handleConn(conn)
	}
}

// todo
func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)

		case msg := <-messages:
			for c := range clients {
				c.ch <- msg
			}
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{conn: conn, name: who, ch: ch}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	if err := input.Err(); err != nil {
		fmt.Println("Error reading from client:", err)
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

// todo
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("error write to client", err)

			return
		}
	}
}
