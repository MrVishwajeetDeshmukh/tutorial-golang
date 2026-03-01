package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func optimizedCommunication() {
	ch := make(chan int, 100000) // Use a buffered channel

	for i := 0; i < 100000; i++ {
		go func(n int) {
			time.Sleep(1 * time.Millisecond) // Artificial delay
			ch <- n
		}(i)
	}

	for i := 0; i < 100000; i++ {
		fmt.Println(<-ch) // Output the received value
	}
}

func main() {
	go optimizedCommunication()
	http.ListenAndServe("localhost:6060", nil)
}
