package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func populateChan(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- rand.Intn(1000)
		time.Sleep(50 * time.Millisecond)
	}
}

func display(ch1, ch2 <-chan int) {
	ch1Open, ch2Open := true, true

	// Loop continuously if at least one of the two channels is open
	for ch1Open || ch2Open {
		select {
		case value, ok := <-ch1:
			if !ok {
				ch1Open = false // Detect ch1 closed
			} else {
				fmt.Printf("ch1: %d\n", value)
			}
		case value, ok := <-ch2:
			if !ok {
				ch2Open = false // Detect ch2 closed
			} else {
				fmt.Printf("ch2: %d\n", value)
			}
		}
	}
	fmt.Println("All channels are closed, terminating display function.")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	channel1 := make(chan int)
	channel2 := make(chan int)

	go display(channel1, channel2)

	wg.Add(2)
	go populateChan(channel1, &wg)
	go populateChan(channel2, &wg)

	// Wait for all populateChan goroutines to finish
	wg.Wait()

	// Data generation completed, close channels to notify display goroutine
	close(channel1)
	close(channel2)

	time.Sleep(300 * time.Millisecond) // Wait briefly for display() function's output
	fmt.Println("main function terminated.")
}
