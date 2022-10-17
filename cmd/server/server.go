package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/schollz/progressbar/v3"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		handleListenerError(err)
	}

	currentNumberOfClients := 0

	updates := make(chan int)
	// Progress bar setup
	bar := progressbar.DefaultBytes(-1)
	go func(updates chan int) {
		for update := range updates {
			bar.Add(update)
		}
	}(updates)
	// Infinite loop to accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			handleAcceptError(err)
		}

		currentNumberOfClients++

		fmt.Println("Current number of TCP clients: ", currentNumberOfClients)
		go handleConn(conn, updates)
	}
}

func handleConn(c net.Conn, updates chan int) {
	defer c.Close()
	for {
		packetSize := rand.Intn(10000 - 1) + 1
		payload := make([]byte, packetSize)
		rand.Read(payload)
		_, err := c.Write(payload)
		// Send update through channel for progress bar
		updates <- packetSize
		// Handle connection write error
		if err != nil {
			handleConnWriteError(err)
			return
		}
	}
}

func handleListenerError(err error) {
	log.Fatal("Error when trying to listen on port 8000:", err)
}

func handleAcceptError(err error) {
	log.Fatal("Error accepting the connection:", err)
}

func handleConnWriteError(err error) {
	log.Println("Error writing to client:", err)
}
