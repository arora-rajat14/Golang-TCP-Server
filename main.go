package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// go run main.go 9090
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <port>")
		os.Exit(1)
	}
	// :9090
	port := fmt.Sprintf(":%s", os.Args[1])

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error occurred while listening to specific port number.")
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Listening on port %s\n", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection, err", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		reader.ReadBytes(byte('\n'))
	}
}
