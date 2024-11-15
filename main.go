package main

import (
	"fmt"
	"net"

	"netcat/functions"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	Connection := functions.NewConnection()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}
	
		go Connection.HandleClient(conn)
	}
}
