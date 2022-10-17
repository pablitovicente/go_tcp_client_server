package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

// Custom Writer
type dataHandler struct {
	clientId int
}

func (dh *dataHandler) Write(data []byte) (int, error) {
	// Do whatever we want with the data
  fmt.Printf("Got %d bytes from conn for clientId %d\n", len(data), dh.clientId)
	// TODO implement a simple protocol interpreter
	return len(data), nil
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(clientId int, wg *sync.WaitGroup) {
			defer wg.Done()
			conn, err := net.Dial("tcp", "localhost:8000")
			if err != nil {
				handleConnError(err)
			}
			defer conn.Close()
			handler := dataHandler{
				clientId: clientId,
			}
			bufio.NewReader(conn).WriteTo(&handler)
		}(i, &wg)
	}

	wg.Wait()
}

func handleConnError(err error) {
	log.Fatal("Error connecting to server:", err)
}
