package main

import (
	"fmt"
	"net"
)

func main() {
	// Attempt to connect to port 80 of google.com using the TCP protocol
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer conn.Close() // Close connection when the program terminates

	fmt.Println("Connection successful:", conn.RemoteAddr())
}
