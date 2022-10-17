package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Custom Writer
type dataHandler struct {
}

func (dh *dataHandler) Write(data []byte) (int, error) {
	// Do whatever we want with the data
  fmt.Printf("Got %d bytes from conn\n", len(data))
	// TODO implement a simple protocol interpreter
	return len(data), nil
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		handleConnError(err)
	}
	defer conn.Close()
	handler := dataHandler{}
	// Create a reader for the connection and then "pipe it" through
	// something that implements the Reader interface
	bufio.NewReader(conn).WriteTo(&handler)
}

func handleConnError(err error) {
	log.Fatal("Error connecting to server:", err)
}
