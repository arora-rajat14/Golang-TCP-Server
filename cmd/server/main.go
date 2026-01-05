package main

import (
	"fmt"
	"os"

	"github.com/arora-rajat14/Golang-TCP-Server/internal/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/server/main.go <port>")
		os.Exit(1)
	}

	port := fmt.Sprintf(":%s", os.Args[1])
	srv := server.NewServer(port)

	if err := srv.Start(); err != nil {
		fmt.Printf("Server failed: %v\n", err)
		os.Exit(1)
	}
}
