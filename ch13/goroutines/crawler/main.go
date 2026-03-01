package main

import (
	"fmt"
	"strings"
	"time"
)

const divideBar = "================================================"

func main() {
	var confirm string
	useGoroutine := false
	fmt.Print("Would you like to use goroutines? (y/N): ")
	fmt.Scanf("%s\n", &confirm)

	if strings.ToUpper(confirm) == "Y" {
		useGoroutine = true
	}

	index, err := Indexing()
	if err != nil {
		panic(err)
	}

	fmt.Println(divideBar)
	fmt.Printf("Found a total of %d news articles.\n", len(index.Link))
	fmt.Println("Fetching news content.")

	start := time.Now()
	err = Crawler(index, useGoroutine)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	fmt.Println("All news content fetched successfully.")
	fmt.Println(divideBar)
	for _, link := range index.Link {
		fmt.Printf("TITLE  : %v\n", link.Title)
		fmt.Printf("LINK   : %v\n", link.URI)
		fmt.Printf("CONTENT: %d characters of content\n", len(link.Content))
		fmt.Println(divideBar)
	}
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
