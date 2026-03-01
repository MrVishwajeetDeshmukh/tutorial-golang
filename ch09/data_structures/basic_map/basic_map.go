package main

import "fmt"

func getRank(score int) string {
	switch {
	case score > 90:
		return "A"
	case score > 80:
		return "B"
	case score > 70:
		return "C"
	default:
		return "F"
	}
}

func main() {
	// The map data structure can be allocated using make(map[keyType]valueType)
	// The key is the index used to get a value
	// The value is what is returned when using the key as an index
	mapScore := make(map[string]int)

	for i := 0; i < 3; i++ {
		var subjectName string
		var subjectScore int

		fmt.Print("Please enter the subject name: ")
		fmt.Scanf("%s\n", &subjectName)

		fmt.Printf("Please enter the score for %s: ", subjectName)
		fmt.Scanf("%d\n", &subjectScore)

		mapScore[subjectName] = subjectScore
	}

	fmt.Print("\n")
	fmt.Print("============================")
	fmt.Print("\n")

	// When using range with a map,
	// it iterates by assigning the key and value to name and score
	for name, score := range mapScore {
		fmt.Printf("The grade for %s is %s!\n", name, getRank(score))
	}
}
