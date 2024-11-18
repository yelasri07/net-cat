package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"TCPChat/functions"
)

var port = ":8989"

func main() {
	// Check for custom port argument; default to ":8989".
	arguments := os.Args
	if len(arguments) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	if len(arguments) == 2 {
		port = ":" + arguments[1]
	}

	// Start listening for incoming connections.
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on the port %s\n", port)
	defer listener.Close()

	// Create a new Connections instance to manage clients and messages.
	connections := functions.NewConnection()

	// Accept and handle client connections in separate goroutines.
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go connections.HandleConnection(conn)
	}
}
