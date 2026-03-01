package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Resolve server address
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed to resolve address:", err)
		return
	}
	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)

	for {
		// Receive message input from user
		fmt.Print("Enter message: ")
		text, _ := reader.ReadString('\n')
		// Send message to server
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Failed to transmit data:", err)
			continue
		}
		// Receive response from server
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Failed to receive response:", err)
			continue
		}
		fmt.Print("Response from server:", string(buffer[:n]))
	}
}
