package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server")
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error listening", err.Error())

		return
	}

	for {

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buff := make([]byte, 512)
		_, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Error ", err.Error())
			return
		}

		fmt.Printf("Received data: %v", string(buff))
	}
}
