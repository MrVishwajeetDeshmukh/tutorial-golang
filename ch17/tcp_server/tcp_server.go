package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Listen for TCP connections on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("TCP server is running on port 8080.")

	for {
		// Accept client connection request
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}
		// Handle each connection concurrently in a goroutine
		go handleConnection(conn)
	}
}

// Client connection handler function
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		// Read message from client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client connection closed:", err)
			return
		}
		fmt.Print("Received message:", message)
		// Eho received message back to client
		_, err = conn.Write([]byte("Server response: " + message))
		if err != nil {
			fmt.Println("Failed to send message:", err)
			return
		}
	}
}
