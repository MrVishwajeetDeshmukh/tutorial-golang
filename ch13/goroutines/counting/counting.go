package main

import (
	"fmt"
	"time"
)

func couting1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("counting1: index %d\n", i)
		time.Sleep(1 * time.Nanosecond)
	}
}

func couting2() {
	for i := 0; i < 5; i++ {
		fmt.Printf("counting2: index %d\n", i)
		time.Sleep(1 * time.Nanosecond)
	}
}

func main() {
	go couting1()
	couting2()
}
