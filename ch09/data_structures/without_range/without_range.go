package main

import "fmt"

func main() {
	score := make([]int, 5)

	// Enter scores for 5 subjects from the user
	for i := 0; i < len(score); i++ {
		fmt.Printf("Enter the score for subject %d: ", i+1)
		fmt.Scanf("%d\n", &score[i])
	}
	fmt.Println("============================")

	// Iterate through all slice elements and calculate the average score
	sumScore := 0
	for i := 0; i < len(score); i++ {
		sumScore += score[i]
	}
	meanScore := float64(sumScore) / float64(len(score))
	fmt.Printf("Your average score is %.2f.\n", meanScore)
}
