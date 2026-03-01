package main

import "fmt"

// getTranslatedWord translates the given English word into the selected language.
// Supported languages:
//   en = English (no change)
//   es = Spanish
//   fr = French
func getTranslatedWord(engWord string, language string) (translatedWord string) {

	// If English, return the original word
	if language == "en" {
		return engWord

	// Spanish translations
	} else if language == "es" {
		if engWord == "apple" {
			return "manzana"
		} else if engWord == "banana" {
			return "plátano"
		} else if engWord == "grape" {
			return "uva"
		} else {
			return "no matching word!"
		}

	// French translations
	} else if language == "fr" {
		if engWord == "apple" {
			return "pomme"
		} else if engWord == "banana" {
			return "banane"
		} else if engWord == "grape" {
			return "raisin"
		} else {
			return "no matching word!"
		}

	// Unsupported language
	} else {
		return "no matching word!"
	}
}

func main() {
	var word, language string

	// Ask for English word
	fmt.Print("Enter the word you want to translate (apple, banana, grape): ")
	fmt.Scanf("%s\n", &word)

	// Ask for target language
	fmt.Print("Enter the target language (supported: es, fr, en): ")
	fmt.Scanf("%s\n", &language)

	// Perform translation
	translatedWord := getTranslatedWord(word, language)

	fmt.Print("\n")
	fmt.Printf("Translated word: %v\n", translatedWord)
}
