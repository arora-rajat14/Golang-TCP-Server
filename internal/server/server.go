package server

import (
	"fmt"
	"net"
)

// Server represents a TCP server.
type Server struct {
	addr string
}

// NewServer creates a new Server instance.
func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

// Start begins listening on the configured address.
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", s.addr, err)
	}
	defer listener.Close()

	fmt.Printf("Listening on port %s\n", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}
