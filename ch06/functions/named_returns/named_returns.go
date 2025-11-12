package main

import (
	"fmt"
	"time"
)

// getAgeDetails calculates a person's age and age range
// based on the year they were born.
func getAgeDetails(bornYear int) (age int, ageRange int) {
	// Get the current year
	currentYear := time.Now().Year()

	// Calculate the age by subtracting the birth year from the current year.
	// Adding +1 since in some cultures (like Korea) age starts at 1 when born.
	ageValue := currentYear - bornYear + 1

	// Determine the age range (e.g., 20s, 30s, 40s, etc.)
	// by dividing the age by 10 and multiplying back by 10.
	// For example, 31 → 30s, 45 → 40s.
	ageRangeValue := int(ageValue/10) * 10

	return ageValue, ageRangeValue
}

func main() {
	// Declare a variable to store the user's birth year
	var bornYear int

	// Ask the user to enter their birth year
	fmt.Print("Enter your birth year: ")
	fmt.Scanf("%d\n", &bornYear)

	// Call getAgeDetails with the user's birth year
	// and store the returned values in age and ageRange.
	age, ageRange := getAgeDetails(bornYear)

	// Print a blank line for better readability
	fmt.Print("\n")

	// Display the calculated age and age range
	fmt.Printf("You are %d years old and in your %ds!\n", age, ageRange)
}
