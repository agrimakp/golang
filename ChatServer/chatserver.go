package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// create a map of client address, message, connection

var clients = make(map[string]net.Conn)
var messages = make(chan message)
var leaving = make(chan message)

type message struct {
	text    string
	address string
}

func newMessage(msg string, conn net.Conn) message {
	addr := conn.RemoteAddr().String()
	return message{
		text:    addr + msg,
		address: addr,
	}
}

func handle(conn net.Conn) {
	clients[conn.RemoteAddr().String()] = conn
	// put or get a message from a channel

	messages <- newMessage(" joined.", conn)
	log.Printf("halo = %v\n", messages)

	input := bufio.NewScanner(conn)

	for input.Scan() {
		messages <- newMessage(": "+input.Text(), conn)
	}
	// Delete client
	delete(clients, conn.RemoteAddr().String())

	leaving <- newMessage("has left.", conn)

	// ignore errors
	conn.Close()
}

type Snake struct {
	x int
	y int
}

func broadcaster() {
	for {
		//lets a goroutine wait on multiple communication operations
		select {
		case msg := <-messages:

			for _, conn := range clients {
				if msg.address == conn.RemoteAddr().String() {
					continue
				}
				// ignore network errors
				fmt.Fprintln(conn, msg.text)
			}

		case msg := <-leaving:
			for _, conn := range clients {

				// ignore network errors
				fmt.Fprintln(conn, msg.text)
			}
		}
	}
}

func main() {

	// creates server and listens for incoming connections on a local network address
	// listen and accept incomming network connections from clients
	listen, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		// raised when an unexpected error occurs
		panic(err)
	}

	go broadcaster()

	//  an infinite loop for accepting connections
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// go starts a goroutine
		go handle(conn)
	}
}
