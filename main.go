package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"TCPChat/functions"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	PORT := ":" + arguments[1]
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on the port %s\n", PORT)
	defer listener.Close()
	connections := functions.NewConnection()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go connections.HandleConnection(conn)
	}
}
