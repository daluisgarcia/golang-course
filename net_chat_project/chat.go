package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string) // Global messages channel
)

var (
	hostS = flag.String("host", "localhost", "The host to scan")
	portS = flag.Int("port", 3090, "The port to scan")
)

func MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	message := make(chan string)
	go MessageWrite(conn, message)

	clientName := conn.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome to the server %s!\n", clientName) // Individual client message channel
	messages <- fmt.Sprintf("%s has joined the server\n", clientName)

	incomingClients <- message

	inputMessage := bufio.NewScanner(conn) // Read messages from the connection
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingClients <- message

	messages <- fmt.Sprintf("%s has left the server\n", clientName)
}

func Broadcast() { // Sends messages to all clients
	clients := make(map[Client]bool)

	for {
		select {
		case msg := <-messages:
			for client := range clients {
				client <- msg
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() { // Server main function for the chat
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *hostS, *portS))

	if err != nil {
		log.Fatal(err)
	}

	go Broadcast()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go HandleConnection(conn)
	}
}
