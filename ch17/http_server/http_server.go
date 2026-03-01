package main

import (
	"fmt"
	"net/http"
)

// Root path request processing function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a web server made with Go!")
}

func main() {
	// Map request path to handler function
	http.HandleFunc("/", helloHandler)
	fmt.Println("HTTP server is running on port 8080.")
	// Start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
