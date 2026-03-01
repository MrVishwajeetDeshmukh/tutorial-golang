package main

import "fmt"

func getSpanishWord(engWord string) (spanWord string) {
	var word string
	if engWord == "apple" {
		word = "manzana"
	} else if engWord == "banana" {
		word = "plátano"
	} else if engWord == "grape" {
		word = "uva"
	} else {
		word = "no matching word!"
	}
	return word
}

func main() {
	var searchWord string

	fmt.Print("Please enter the word you want to look up (apple, banana, grape): ")
	fmt.Scanf("%s\n", &searchWord)

	fmt.Print("\n")
	spanWord := getSpanishWord(searchWord)
	fmt.Printf("Found word: %v\n", spanWord)
}
