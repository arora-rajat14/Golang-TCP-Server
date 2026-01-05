package server

import (
	"net"
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {
	// Pick a random port by letting the OS verify it's free, but our server relies on string addr.
	// We'll trust that 9091 is free for this test or let it fail if not.
	// A better way is to bind to separate port 0 and let net.Listen pick,
	// but our current implementation takes a fixed string.
	addr := ":9091"
	srv := NewServer(addr)

	// Run server in a goroutine
	go func() {
		if err := srv.Start(); err != nil {
			// This might log if the server is closed during test, which is expected
			t.Logf("Server stopped: %v", err)
		}
	}()

	// Give it a moment to start
	time.Sleep(100 * time.Millisecond)

	// Try to connect
	conn, err := net.Dial("tcp", "localhost"+addr)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Send a message
	msg := "Hello Server\n"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		t.Fatalf("Failed to write to server: %v", err)
	}

	// Read response (Echo: ...)
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("Failed to read from server: %v", err)
	}

	expected := "Echo: " + msg
	if string(buf[:n]) != expected {
		t.Errorf("Expected response %q, got %q", expected, string(buf[:n]))
	}
}
