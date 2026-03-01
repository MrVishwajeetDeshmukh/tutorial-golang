package main

import (
	"fmt"
	"net"
)

func main() {
	// Create UDP address
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Failed to resolve address:", err)
		return
	}
	// Listen for UDP connections
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Failed to start UDP server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("UDP server is running on port 8080.")

	buffer := make([]byte, 1024)
	for {
		// Receive data from client
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Failed to receive data:", err)
			continue
		}
		message := string(buffer[:n])
		fmt.Printf("Received message: %s\n", message)
		// Send response to client
		_, err = conn.WriteToUDP([]byte("Server response: "+message), clientAddr)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
