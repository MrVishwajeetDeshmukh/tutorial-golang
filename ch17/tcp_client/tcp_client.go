package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Attempt TCP connection to localhost port 8080
	conn, err := net.Dial("tcp", "localhost:8080")
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
		_, err = fmt.Fprintf(conn, text)
		if err != nil {
			fmt.Println("Failed to send message:", err)
			return
		}
		// Receive response from server
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Failed to receive server response:", err)
			return
		}
		fmt.Print("Response from server:", message)
	}
}
