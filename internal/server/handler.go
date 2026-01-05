package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// handleConnection manages an individual client connection.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		// Read up from the connection until we hit a newline
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Failed to read data, err: %v\n", err)
			}
			return
		}

		fmt.Printf("request: %s", bytes)
		line := fmt.Sprintf("Echo: %s", bytes)
		fmt.Printf("Response: %s", line)

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Printf("Failed to write data, err: %v\n", err)
			return
		}
	}
}
