package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

func main() {
	// Create trace file
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Start trace
	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	// Simulated task
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
