package main

import "fmt"

// getSpanishWord returns the Spanish translation of the English word.
func getSpanishWord(engWord string) (spanishWord string) {
	var word string
	switch engWord {
	case "apple":
		word = "manzana"
	case "banana":
		word = "plátano"
	case "grape":
		word = "uva"
	default:
		word = "no matching word!"
	}
	return word
}

func main() {
	var searchWord string

	fmt.Print("Enter the word you want to translate (apple, banana, grape): ")
	fmt.Scanf("%s\n", &searchWord)

	fmt.Print("\n")
	spanishWord := getSpanishWord(searchWord)
	fmt.Printf("Translated word: %v\n", spanishWord)
}
