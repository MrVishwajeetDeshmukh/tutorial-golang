package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func increment() {
	mu.Lock()
	defer mu.Unlock() // Call Unlock with defer right after Lock
	count++
	// Even if a panic occurs here, the defer Unlock on line 15 will be executed
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("Count:", count)
}
