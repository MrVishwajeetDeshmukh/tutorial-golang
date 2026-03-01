package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Send HTTP GET request
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}
	fmt.Println("Response status code:", resp.StatusCode)
	fmt.Println("Response body length:", len(body))
}
