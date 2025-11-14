package main

import "fmt"

// getTranslatedWord takes an English word and a target language code,
// then returns the translated version.
// Supported languages:
//   en = English (no change)
//   ko = Korean
//   jp = Japanese
// If no match is found, it returns a fallback message.
func getTranslatedWord(
	engWord string,
	language string) (translatedWord string) {

	// If the user wants English, just return the original word.
	if language == "en" {
		return engWord

	// apple translations
	} else if engWord == "apple" && language == "ko" {
		return "사과"
	} else if engWord == "apple" && language == "jp" {
		return "林檎"

	// banana translations
	} else if engWord == "banana" && language == "ko" {
		return "바나나"
	} else if engWord == "banana" && language == "jp" {
		return "バナナ"

	// grape translations
	} else if engWord == "grape" && language == "ko" {
		return "포도"
	} else if engWord == "grape" && language == "jp" {
		return "葡萄"

	// If nothing matches
	} else {
		return "no matching word!"
	}
}

func main() {
	var word, language string

	// Ask the user which word they want to translate
	fmt.Print("Enter the word you want to translate (apple, banana, grape): ")
	fmt.Scanf("%s\n", &word)

	// Ask the user which language they want to translate to
	fmt.Print("Enter the target language (supported: ko, jp, en): ")
	fmt.Scanf("%s\n", &language)

	// Get translated result
	translatedWord := getTranslatedWord(word, language)

	fmt.Print("\n")
	fmt.Printf("Translated word: %v\n", translatedWord)
}
